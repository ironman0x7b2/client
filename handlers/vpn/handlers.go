package vpn

import (
	"log"
	"net/http"
	"os/exec"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/handlers/errors"
	"github.com/ironman0x7b2/client/utils"
)

const MODULE = "vpn"

func connectVPNHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newVPN(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errors.ErrorParseRequestBody(MODULE))
			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, errors.ErrorValidateRequestBody(MODULE))
			log.Println(err.Error())
			return
		}

		var subscriptionID string
		var port string

		if body.SubscriptionID == "" {
			port = strconv.FormatUint(body.NodePort, 10)
			subscriptionID, err = utils.StartSubscription(cli, body.From, body.FromAddress, body.Password,
				body.NodeID, body.ResolverID, body.NodeIP, port, body.Amount, body.Memo, body.Fees,
				body.GasPrices, body.Gas, body.GasAdjustment, w)

			if err != nil {
				log.Println(err.Error())
				return
			}
			utils.ConnectVPN(body.From, body.Password, subscriptionID, body.NodeIP, port)
		}
		utils.ConnectVPN(body.From, body.Password, body.SubscriptionID, body.NodeIP, port)
	}
}

func endVPNHandler() http.HandlerFunc {
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

func getSubscriptionsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		address, err := sdk.AccAddressFromHex(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errors.ErrorDecodeAddress(MODULE))
			return
		}

		subscriptions, _err := cli.GetSubscriptonsOfClientFromRPC(address, MODULE)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, _err)
			return
		}

		utils.WriteResultToResponse(w, 200, subscriptions)
	}
}
