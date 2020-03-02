package vpn

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	hub "github.com/sentinel-official/hub/types"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sentinel-official/dvpn-node/node"
	nodeTypes "github.com/sentinel-official/dvpn-node/types"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func startSubscriptionHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newSubscription(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate request body",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		msg, err := messages.NewSubscription(body.FromAddress, body.Amount, body.NodeID, body.ResolverID).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the withdrawRewards message",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to broadcast the transaction",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		log.Println("Start subscription transaction compleated with hash  %s", res.TxHash)

		utils.WriteResultToResponse(w, 200, res.TxHash)
	}
}

func submitTxHashToNodeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newsubmitTxHashToNode(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate request body",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		strPort := strconv.FormatUint(uint64(body.NodePort), 10)
		url := "https://" + body.NodeIP + ":" + strPort + "/subscriptions"
		message := map[string]interface{}{
			"tx_hash": body.TxHash,
		}

		bz, err := json.Marshal(message)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to broadcast the transaction",
				Info:    err.Error(),
			})

			log.Println(err)
			return
		}

		resp, err := client.Post(url, "application/json", bytes.NewBuffer(bz))
		if err != nil {
			log.Fatalln(err.Error())
		}

		if resp.StatusCode != 201 {
			log.Fatalln("Error while submitting tx_hash to node")
		}

		_body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("Error while reading response body from node")
		}

		var _resp types.Response
		err = json.Unmarshal(_body, &_resp)
		if err != nil {
			log.Fatalln("Error while unmarshal node response")
		}

		utils.WriteResultToResponse(w, 200, _resp)
	}
}

func getVPNKeyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		body, err := newGetPVNKey(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate request body",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		strPort := strconv.FormatUint(body.NodePort, 10)
		url := "https://" + body.NodeIP + ":" + strPort + "/subscriptions/" + vars["id"] + "/key"
		resp, err := client.Post(url, "application/json", nil)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if resp.StatusCode != 200 {
			log.Fatalln("Error while getting key from node")
		}

		_body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("Error while reading response body from node")
		}

		var _resp types.Response
		err = json.Unmarshal(_body, &_resp)
		if err != nil {
			log.Fatalln("Error while unmarshal node response")
		}

		utils.WriteResultToResponse(w, 200, _resp)
	}
}

func connectVPNHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		body, err := newConnectVPN(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate request body",
				Info:    err.Error(),
			})

			log.Println(err.Error())
			return
		}

		err = ioutil.WriteFile(types.DefaultConfigDir+"/base64.txt", []byte(body.Key), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}

		err = ioutil.WriteFile(types.DefaultConfigDir+"/vpn.ovpn", nil, 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}

		_cmd := "cat " + types.DefaultConfigDir + "/base64.txt" + " | base64 -d > " + types.DefaultConfigDir + "/vpn.ovpn"

		fmt.Println(_cmd)
		cmd := exec.Command("/bin/sh", "-c", _cmd)

		_, err = cmd.Output()
		if err != nil {
			panic(err)
		}

		connCmd := "sudo openvpn " + types.DefaultConfigDir + "/vpn.ovpn "
		go func() {
			cmd = exec.Command("/bin/sh", "-c", connCmd)
			_, err = cmd.Output()
			if err != nil {
				panic(err)
			}
		}()

		strPort := strconv.FormatUint(body.NodePort, 10)
		addr := body.NodeIP + ":" + strPort
		path := "/subscriptions/" + vars["id"] + "/websocket"
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
		defer c.Close()

		kb, err := keys.NewKeyBaseFromDir(types.DefaultConfigDir)
		if err != nil {
			panic(err)
		}

		keyInfo, err := kb.Get(body.AccountName)
		if err != nil {
			log.Println("failed to get the key info from key base", err)
			panic(err)
		}

		var msg nodeTypes.Msg
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

				if err := json.Unmarshal(message, &msg); err == nil {
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
					if err := json.Unmarshal(msg.Data, &_msg); err != nil {
						log.Println(err)
						return
					}

					data := hub.NewBandwidthSignatureData(_msg.ID, _msg.Index, _msg.Bandwidth).Bytes()
					keyInfo.GetPubKey()

					_sign, _, _ := kb.Sign(body.AccountName, body.Password, data)

					msg := node.NewMsgBandwidthSignature(_msg.ID, _msg.Index, _msg.Bandwidth, _msg.NodeOwnerSignature, _sign)

					err = c.WriteMessage(websocket.TextMessage, msg.Bytes())
					if err != nil {
						log.Println("write:", err)
						return
					}
				}

			case <-interrupt:
				log.Println("interrupt")
				err := cmd.Process.Kill()
				if err != nil {
					log.Println("Kill command error", err)
					return
				}
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

		utils.WriteResultToResponse(w, 200, "Open vpn writed")
	}
}

func getSubscriptionsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		address, err := sdk.AccAddressFromHex(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to convert account address",
				Info:    err.Error(),
			})
			return
		}

		subscriptions, _err := cli.GetSubscriptonsOfClientFromRPC(address)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, _err)
			return
		}

		utils.WriteResultToResponse(w, 200, subscriptions)
	}
}
