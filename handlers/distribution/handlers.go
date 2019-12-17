package distribution

import (
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"github.com/tendermint/tendermint/libs/common"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

/**
 * @api {post} /accounts/withdraw-withdrawRewards/{validatorAddress} withdraw withdrawRewards
 * @apiDescription Used to withdraw delegation WithdrawRewards from single validator
 * @apiName Withdraw-withdrawRewards
 * @apiGroup distribution
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"gas":210000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func withdrawRewardsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		body, err := newRewards(r)
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

		msg, err := messages.NewWithdrawRewards(body.FromAddress, vars["validatorAddress"]).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the withdrawRewards message",
				Info:    err.Error(),
			})
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
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}

/**
 * @api {post} /accounts/withdraw-all-withdrawRewards withdraw all withdrawRewards
 * @apiDescription Used to withdraw delegation WithdrawRewards form all validators
 * @apiName Withdraw-all-withdrawRewards
 * @apiGroup distribution
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"gas":210000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func withdrawAllRewardsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newRewards(r)
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

		from, err := sdk.AccAddressFromHex(body.FromAddress)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to convert account address",
				Info:    err.Error(),
			})
			return
		}

		validators, _err := cli.GetDelegatorValidators(from)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, _err)
			return
		}

		var msgs []sdk.Msg
		for _, validator := range validators {
			msg, err := messages.NewWithdrawRewards(body.FromAddress, common.HexBytes(validator.OperatorAddress.Bytes()).String()).Raw()
			if err != nil {
				utils.WriteErrorToResponse(w, 400, &types.Error{
					Message: "failed to prepare the withdrawRewards message",
					Info:    err.Error(),
				})
				return
			}
			msgs = append(msgs, msg)
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, err := cli.Tx(msgs, body.Memo, body.Gas, body.GasAdjustment,
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
