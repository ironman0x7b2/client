package vpn

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkUtils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/handlers/common"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

const MODULE = "vpn"

var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func getResolversNodesHandler(cli *_cli.CLI, cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nodes, _err := cli.GetResolversNodes(cfg, client)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, errorFailedToGetNodes())

			log.Println(_err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, nodes)
	}
}

func getResolverNodesHandler(cli *_cli.CLI, cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		nodes, _err := cli.GetResolverNodes(cfg, client, vars["id"])
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, errorFailedToGetNodes())

			log.Println(_err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, nodes)
	}
}

/**
 * @api {get} /subscriptions/{address} get subscriptions of user
 * @apiDescription Used to get user subscription details
 * @apiName GetSubscriptions
 * @apiGroup vpn
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getSubscriptionsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		address, err := sdk.AccAddressFromHex(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorDecodeAddress(MODULE))
			return
		}

		subscriptions, _err := cli.GetSubscriptonsOfClientFromRPC(address)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, errorFailedToGetSubscriptions())

			log.Println(_err.Error())
			return
		}
		if len(subscriptions) == 0 {
			utils.WriteErrorToResponse(w, 400, errorNoSubscriptions())

			log.Println(errMsgNoSubscriptions)
			return
		}

		utils.WriteResultToResponse(w, 200, subscriptions)
	}
}

/**
 * @api {post} /subscription to start subscription
 * @apiDescription Used to start subscription for vpn connection
 * @apiName StartSubscription
 * @apiGroup vpn
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"amount":[{"denom":"tsent","value":10}],
 *	"node_id":"node1",
 *	"resolver_id":"reso0",
 *	"node_ip":"127.0.0.1",
 *	"node_port":8080,
 *	"gas":2100000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func startSubscriptionHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newSubscription(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorParseRequestBody(MODULE))
			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorValidateRequestBody(MODULE))
			log.Println(err.Error())
			return
		}

		msg, err := messages.NewSubscription(body.FromAddress, body.Amount, body.NodeID, body.ResolverID).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToPrepareMsg(MODULE))

			log.Println(err.Error())
			return
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, _err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToBroadcastTransaction(MODULE))

			log.Println(_err.Error())
			return
		}

		time.Sleep(4 * time.Second)
		data, err := sdkUtils.QueryTx(cli.CLIContext, res.TxHash)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errorQuerySubscriptionTransaction())
			log.Println(err)
			return
		}

		subscriptionID := data.Events[1].Attributes[0].Value
		log.Println("Start _newVPN transaction completed with hash  and subsription_id", res.TxHash, subscriptionID)

		port := strconv.FormatUint(body.NodePort, 10)
		_url := "https://" + body.NodeIP + ":" + port + "/subscriptions"

		message := map[string]interface{}{
			"tx_hash": res.TxHash,
		}

		bz, err := json.Marshal(message)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errorMarshallNodeRequestBody())

			log.Println(err)
			return
		}

		resp, err := client.Post(_url, "application/json", bytes.NewBuffer(bz))
		if err != nil {

			log.Println(err.Error())
			return
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

		var _res struct {
			TxHash         string `json:"tx_hash"`
			SubscriptionID string `json:"subscription_id"`
		}

		_res.TxHash = res.TxHash
		_res.SubscriptionID = subscriptionID

		utils.WriteResultToResponse(w, 200, _res)
	}
}

/**
 * @api {post} /vpn to start vpn connection
 * @apiDescription Used to start vpn connection
 * @apiName StartVPNConnection
 * @apiGroup vpn
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"node_id":"node1",
 *	"resolver_id":"reso0",
 *	"node_ip":"127.0.0.1",
 *	"node_port":8080,
 *	"gas":2100000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func startVPNConnectionHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newVPN(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorParseRequestBody(MODULE))
			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorValidateRequestBody(MODULE))
			log.Println(err.Error())
			return
		}

		address, err := sdk.AccAddressFromHex(body.FromAddress)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorDecodeAddress(MODULE))
			return
		}

		subscriptions, _err := cli.GetSubscriptonsOfClientFromRPC(address)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, errorFailedToGetSubscriptions())

			log.Println(_err.Error())
			return
		}
		if len(subscriptions) == 0 {
			utils.WriteErrorToResponse(w, 400, errorNoSubscriptions())

			log.Println(errMsgNoSubscriptions)
			return
		}

		port := strconv.FormatUint(body.NodePort, 10)
		subscriptionID, err := utils.GetVPNConnectionSubscription(subscriptions, body.ResolverID, body.NodeID)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errorNoSubscriptions())

			log.Println(errMsgNoSubscriptions)
			return
		}

		utils.ConnectVPN(client, body.From, body.Password, subscriptionID, body.NodeIP, port)
	}
}

/**
 * @api {delete} /vpn to end vpn connection
 * @apiDescription Used to end vpn connection
 * @apiName EndVPNConnection
 * @apiGroup vpn
 */

func endVPNConnectionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		disConnCmd := "sudo killall openvpn"

		cmd := exec.Command("/bin/sh", "-c", disConnCmd)
		_, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		utils.WriteResultToResponse(w, 200, "connection ended")
	}
}
