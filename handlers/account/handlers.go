package account

import (
	"log"
	"net/http"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	
	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/models"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

/**
 * @api {get} /accounts/{address} get account
 * @apiDescription Used to get account details
 * @apiName GetAccount
 * @apiGroup account
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getAccountHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		
		address, err := sdk.AccAddressFromHex(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to decode the address",
				Info:    err.Error(),
			})
			
			log.Println(err.Error())
			return
		}
		
		account, err := cli.GetAccount(address)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to query the account",
				Info:    err.Error(),
			})
			
			log.Println(err.Error())
			return
		}
		
		_account := models.NewAccountFromRaw(account)
		utils.WriteResultToResponse(w, 200, _account)
	}
}

/**
 * @api {post} /transfer transfer coins
 * @apiDescription Used to transfer coins
 * @apiName TransferCoins
 * @apiGroup account
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"to_address":"35BC67ABA8E19D9462F2C9CEA15AC8643E77166F",
 *	"amount":[{"denom":"tsent","value":10}],
 *	"password":"password",
 *	"gas":210000
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func transferCoinsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newTransferCoins(r)
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
		
		msg, err := messages.NewSend(body.FromAddress, body.ToAddress, body.Amount).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the transfer message",
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

/**
 * @api {get} /accounts/{address}/delegations get delegator delegations
 * @apiDescription Used to get all delegations of delegator
 * @apiName GetDelegatorDelegations
 * @apiGroup account
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getDelegatorDelegationsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		
		address, err := sdk.AccAddressFromHex(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to get address",
				Info:    err.Error(),
			})
			
			log.Println(err.Error())
			return
		}
		
		delegations, _err := cli.GetDelegatorDelegations(address)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, _err)
			
			log.Println(_err.Info)
			return
		}
		
		_delegations := models.NewDelegationsFromRaw(delegations)
		utils.WriteResultToResponse(w, 200, _delegations)
	}
}

/**
 * @api {get} /accounts/{address}/delegations/validators get delegator validators
 * @apiDescription Used to get all validators of delegator
 * @apiName GetDelegatorValidators
 * @apiGroup account
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getDelegatorValidatorsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		
		validators, _err := cli.GetDelegatorValidators(vars["address"])
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, _err)
			
			log.Println(_err.Info)
			return
		}
		
		utils.WriteResultToResponse(w, 200, validators)
	}
}
