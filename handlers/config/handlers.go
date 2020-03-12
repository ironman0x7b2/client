package config

import (
	"log"
	"net/http"

	"github.com/ironman0x7b2/client/handlers/common"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

/**
 * @api {get} /config get config
 * @apiDescription Used to get config details
 * @apiName GetConfig
 * @apiGroup config
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

const MODULE = "config"

func getConfigHandler(config *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.WriteResultToResponse(w, 200, config)
	}
}

/**
 * @api {put} /config update config
 * @apiDescription Used to update config details
 * @apiName UpdateConfig
 * @apiGroup config
 * @apiParamExample {json} Request-Example:
 * {
 *	"chain_id":"sentinel-turing-2",
 *	"rpc_address":"ip:port",
 *  "explorer_address":"ip:port",
 *  "verifier_dir": "/home/user/.sentinel/client",
 * "keys_dir": "/home/user/.sentinel/client",
 * "resolver_address": "ip:port",
 * "trust_node": false,
 * "kill_switch": false
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func updateConfigHandler(config *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newUpdateConfig(r)
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

		updates := &types.Config{
			ChainID:         body.ChainID,
			RPCAddress:      body.RPCAddress,
			ExplorerAddress: body.ExplorerAddress,
			VerifierDir:     body.VerifierDir,
			TrustNode:       body.TrustNode,
			KeysDir:         body.KeysDir,
			Resolvers:       body.Resolvers,
			KillSwitch:      body.KillSwitch,
		}

		if err := config.UpdateHook(updates); err != nil {
			utils.WriteErrorToResponse(w, 500, errorFailedToCallUpdateHook())

			log.Println(err.Error())
			return
		}

		config.Update(updates)
		if err := config.SaveToPath(""); err != nil {
			utils.WriteErrorToResponse(w, 500, errorFailedToSaveConfig())

			log.Println(err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, config)
	}
}
