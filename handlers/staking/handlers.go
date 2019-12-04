package staking

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

/**
 * @api {get} /accounts/{address}/delegations get Delegator Delegations
 * @apiDescription get Delegator Delegations
 * @apiName getDelegatorDelegations
 * @apiGroup staking
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
			return
		}

		delegations, _err := cli.GetDelegatorDelegations(address)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, _err)
			return
		}

		_delegations := models.NewDelegationsFromRaw(delegations)
		utils.WriteResultToResponse(w, 200, _delegations)
	}
}

/**
 * @api {get} /accounts/{address}/delegations/validators get Delegator Validators
 * @apiDescription get Delegator Validators
 * @apiName GetDelegatorValidators
 * @apiGroup staking
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getDelegatorValidatorsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		address, err := sdk.AccAddressFromHex(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to get address",
				Info:    err.Error(),
			})
			return
		}

		validators, _err := cli.GetDelegatorValidators(address)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, _err)
			return
		}

		_validators := models.NewValidatorsFromRaw(validators)
		utils.WriteResultToResponse(w, 200, _validators)
	}
}

/**
 * @api {get} /validators get All Validators
 * @apiDescription get All Validators
 * @apiName GetAllValidators
 * @apiGroup staking
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getAllValidatorsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validator, err := cli.GetAllValidators()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)
			return
		}

		utils.WriteResultToResponse(w, 200, validator)
	}
}

/**
 * @api {get} /validators/{address} get Validators
 * @apiDescription get Validators
 * @apiName GetValidators
 * @apiGroup staking
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getValidatorHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		validator, err := cli.GetValidator(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)
			return
		}

		utils.WriteResultToResponse(w, 200, validator)
	}
}

/**
 * @api {post} delegations/{validatorAddress} delegate Coins
 * @apiDescription delegate Coins
 * @apiName Delegate
 * @apiGroup staking
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"amount":[{"denom":"tsent","value":10}],
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

		msg, err := messages.NewDelegate(body.FromAddress, vars["validatorAddress"], body.Amount).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the unbond message",
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
 * @api {put} /re-delegation/{valSrcAddress} redelegate Coins
 * @apiDescription redelegate Coins
 * @apiName ReDelegate
 * @apiGroup staking
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *  "val_dest_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"amount":[{"denom":"tsent","value":10}],
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

		msg, err := messages.NewReDelegate(body.FromAddress, vars["valSrcAddress"], body.ValDestAddress, body.Amount).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the re-delegate message",
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
 * @api {delete} /delegation/{validatorAddress} unbond Coins
 * @apiDescription unbond Coins
 * @apiName Unbond
 * @apiGroup staking
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"amount":[{"denom":"tsent","value":10}],
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

		msg, err := messages.NewUnbond(body.FromAddress, vars["validatorAddress"], body.Amount).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the unbond message",
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
