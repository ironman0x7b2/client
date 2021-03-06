package distribution

import (
	"log"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"github.com/tendermint/tendermint/libs/common"

	_cli "github.com/ironman0x7b2/client/cli"
	_common "github.com/ironman0x7b2/client/handlers/common"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/utils"
)

/**
 * @api {post} /accounts/withdraw-rewards/{validatorAddress} withdraw rewards
 * @apiDescription Used to withdraw delegation WithdrawRewards from single validator
 * @apiName Withdraw-Rewards
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
const MODULE = "distribution"

func withdrawRewardsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		body, err := newRewards(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, _common.ErrorParseRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, _common.ErrorValidateRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		msg, err := messages.NewWithdrawRewards(body.FromAddress, vars["validatorAddress"]).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, _common.ErrorFailedToPrepareMsg(MODULE))

			log.Println(err.Error())
			return
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, _err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, _common.ErrorFailedToBroadcastTransaction(MODULE))

			log.Println(_err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}

/**
 * @api {post} /accounts/withdraw-all-rewards withdraw all Rewards
 * @apiDescription Used to withdraw delegation WithdrawRewards form all validators
 * @apiName Withdraw-all-Rewards
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
			utils.WriteErrorToResponse(w, 400, _common.ErrorParseRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, _common.ErrorValidateRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		from, err := sdk.AccAddressFromHex(body.FromAddress)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, _common.ErrorDecodeAddress(MODULE))

			log.Println(err.Error())
			return
		}

		validators, _err := cli.GetDelegatorValidatorsFromRPC(from)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, _common.ErrorFailedToGetDelegatorValidators(MODULE))

			log.Println(err.Error())
			return
		}

		var msgs []sdk.Msg
		for _, validator := range validators {
			msg, err := messages.NewWithdrawRewards(body.FromAddress, common.HexBytes(validator.OperatorAddress.Bytes()).String()).Raw()
			if err != nil {
				utils.WriteErrorToResponse(w, 400, _common.ErrorFailedToPrepareMsg(MODULE))

				log.Println(err.Error())
				return
			}
			msgs = append(msgs, msg)
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, _err := cli.Tx(msgs, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, _common.ErrorFailedToBroadcastTransaction(MODULE))

			log.Println(_err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}
