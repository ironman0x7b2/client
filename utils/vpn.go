package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkUtils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/websocket"
	"github.com/sentinel-official/dvpn-node/node"
	nodeTypes "github.com/sentinel-official/dvpn-node/types"
	"github.com/sentinel-official/dvpn-node/utils"
	hub "github.com/sentinel-official/hub/types"

	"github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/handlers/errors"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/types"
)

var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func StartSubscription(cli *cli.CLI, from, fromAddress, password, nodeID, resolverID, nodeIP, nodePort string,
	amount types.Coin, memo string, fees types.Coins, gasPrices types.DecCoins, gas uint64, gasAdjustment float64, w http.ResponseWriter) (string, error) {
	msg, err := messages.NewSubscription(fromAddress, amount, nodeID, resolverID).Raw()
	if err != nil {
		utils.WriteErrorToResponse(w, 400, err)
		log.Println(err.Error())
		return "", err
	}

	cli.CLIContext = cli.WithFromName(from)

	res, _err := cli.Tx([]sdk.Msg{msg}, memo, gas, gasAdjustment,
		gasPrices.Raw(), fees.Raw(), password)
	if _err != nil {
		utils.WriteErrorToResponse(w, 400, _err)
		log.Println(_err.Message)
		return "", err
	}

	time.Sleep(4 * time.Second)
	data, err := sdkUtils.QueryTx(cli.CLIContext, res.TxHash)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var _res struct {
		TxHash         string `json:"tx_hash"`
		SubscriptionID string `json:"subscription_id"`
	}

	_res.TxHash = res.TxHash
	subscriptionID := data.Events[1].Attributes[0].Value
	log.Println("Start _newVPN transaction completed with hash  and subsription_id", res.TxHash, subscriptionID)

	_url := "https://" + nodeIP + ":" + nodePort + "/subscriptions"

	message := map[string]interface{}{
		"tx_hash": _res.TxHash,
	}

	bz, err := json.Marshal(message)
	if err != nil {
		utils.WriteErrorToResponse(w, 400, errors.ErrorFailedToMarshallNodeReqMsg())
		log.Println(err)
		return "", err
	}

	resp, err := client.Post(_url, "application/json", bytes.NewBuffer(bz))
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	if resp.StatusCode != 201 {
		log.Println("Error while submitting tx_hash to node")
	}

	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading response body from node")
	}

	var _resp types.Response
	err = json.Unmarshal(_body, &_resp)
	if err != nil {
		log.Println("Error while unmarshal node response")
	}

	return subscriptionID, nil
}

func ConnectVPN(from, password, subscriptionID, nodeIP string, nodePort string) {
	_url := "https://" + nodeIP + ":" + nodePort + "/subscriptions/" + subscriptionID + "/key"
	resp, err := client.Post(_url, "application/json", nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if resp.StatusCode != 200 {
		log.Println("Error while getting key from node")
	}

	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading response body from node")
	}

	var _resp types.Response
	err = json.Unmarshal(_body, &_resp)
	if err != nil {
		log.Println("Error while unmarshal node response")
	}

	decoded, err := base64.StdEncoding.DecodeString(fmt.Sprintf("%v", _resp.Result))

	if err != nil {
		return
	}

	message := map[string]interface{}{
		"signature": "",
	}

	bz, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	_url = "https://" + nodeIP + ":" + nodePort + "/subscriptions/" + subscriptionID + "/sessions"
	resp, err = client.Post(_url, "application/json", bytes.NewBuffer(bz))
	if err != nil {
		return
	}

	if resp.StatusCode != 201 {
		log.Println("Error while init session")
	}

	_body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading response body from node")
	}

	err = json.Unmarshal(_body, &_resp)
	if err != nil {
		log.Println("Error while unmarshal node response")
	}

	err = ioutil.WriteFile(types.DefaultConfigDir+"/vpn.ovpn", decoded, 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	connCmd := "sudo openvpn " + types.DefaultConfigDir + "/vpn.ovpn "
	var cmd *exec.Cmd
	go func() {
		cmd = exec.Command("/bin/sh", "-c", connCmd)
		_, err = cmd.Output()
		if err != nil {
			panic(err)
		}
	}()

	addr := nodeIP + ":" + nodePort
	path := "/subscriptions/" + subscriptionID + "/websocket"
	u := url.URL{Scheme: "wss", Host: addr, Path: path,}
	log.Printf("connecting to %s", u.String())

	dailer := websocket.DefaultDialer
	dailer.ReadBufferSize = 1024
	dailer.WriteBufferSize = 1024

	dailer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c, _, err := dailer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err.Error())
	}

	kb, err := keys.NewKeyBaseFromDir(types.DefaultConfigDir)
	if err != nil {
		panic(err)
	}

	keyInfo, err := kb.Get(from)
	if err != nil {
		log.Println("failed to get the key info from key base", err)
		panic(err)
	}

	var _message nodeTypes.Msg
	var _msg node.MsgBandwidthSignature
	var f bool
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			log.Printf("reading")
			log.Printf("recv: %s", message)

			if err := json.Unmarshal(message, &_message); err == nil {
				f = true
			}
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:

			if f == true {
				if err := json.Unmarshal(_message.Data, &_msg); err != nil {
					log.Println(err)
					return
				}

				data := hub.NewBandwidthSignatureData(_msg.ID, _msg.Index, _msg.Bandwidth).Bytes()
				keyInfo.GetPubKey()

				_sign, _, _ := kb.Sign(from, password, data)

				msg := node.NewMsgBandwidthSignature(_msg.ID, _msg.Index, _msg.Bandwidth, _msg.NodeOwnerSignature, _sign)

				err = c.WriteMessage(websocket.TextMessage, msg.Bytes())
				if err != nil {
					log.Println("write:", err)
					return
				}
			}

		case <-interrupt:
			log.Println("interrupt")
			err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
