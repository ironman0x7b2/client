package account

import (
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/models"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

func getAccountHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		address, err := sdk.AccAddressFromHex(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to decode the address",
				Info:    err.Error(),
			})
			return
		}

		account, err := cli.GetAccount(address)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to query the account",
				Info:    err.Error(),
			})
			return
		}

		_account := models.NewAccountFromRaw(account)
		utils.WriteResultToResponse(w, 200, _account)
	}
}

func transferCoinsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newTransferCoins(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate request body",
				Info:    err.Error(),
			})
			return
		}

		msg, err := messages.NewSend(body.FromAddress, body.ToAddress, body.Amount).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the transfer message",
				Info:    err.Error(),
			})
			return
		}

		res, err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to broadcast the transaction",
				Info:    err.Error(),
			})
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}
