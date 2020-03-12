package staking

import (
	"log"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/handlers/common"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/utils"
)

/**
 * @api {get} /validators get all validators
 * @apiDescription Used to get all validators
 * @apiName GetAllValidators
 * @apiGroup staking
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

const MODULE = "staking"

func getAllValidatorsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validator, err := cli.GetAllValidators(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToGetValidators(MODULE))

			log.Println(err)
			return
		}

		utils.WriteResultToResponse(w, 200, validator)
	}
}

/**
 * @api {get} /validators/{address} get validator
 * @apiDescription Used to get validator details
 * @apiName GetValidator
 * @apiGroup staking
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getValidatorHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		validator, err := cli.GetValidator(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToGetValidator(MODULE))

			log.Println(err)
			return
		}

		utils.WriteResultToResponse(w, 200, validator)
	}
}

/**
 * @api {post} delegations/{validatorAddress} delegate coins
 * @apiDescription Used to delegate coins to validator
 * @apiName Delegate
 * @apiGroup staking
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"amount":{"denom":"tsent","value":10},
 *	"gas":210000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func delegationHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		body, err := newDelegate(r)
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

		msg, err := messages.NewDelegate(body.FromAddress, vars["validatorAddress"], body.Amount).Raw()
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

		utils.WriteResultToResponse(w, 200, res)
	}
}

/**
 * @api {put} /delegation/{valSrcAddress} redelegate coins
 * @apiDescription Used to redelegate coins from one validator to other validator
 * @apiName ReDelegate
 * @apiGroup staking
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *  "val_dest_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"amount":{"denom":"tsent","value":10},
 *	"gas":210000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func reDelegationHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		body, err := newReDelegation(r)
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

		msg, err := messages.NewReDelegate(body.FromAddress, vars["valSrcAddress"], body.ValDestAddress, body.Amount).Raw()
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

		utils.WriteResultToResponse(w, 200, res)
	}
}

/**
 * @api {delete} /delegation/{validatorAddress} unbond coins
 * @apiDescription Used to unbond(undelegate) coins from validator
 * @apiName Unbond
 * @apiGroup staking
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"amount":{"denom":"tsent","value":10},
 *	"gas":210000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func unbondHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		body, err := newUnbond(r)
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

		msg, err := messages.NewUnbond(body.FromAddress, vars["validatorAddress"], body.Amount).Raw()
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

		utils.WriteResultToResponse(w, 200, res)
	}
}
