package vpn

import (
	"log"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

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

		utils.WriteResultToResponse(w, 200, res)
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
