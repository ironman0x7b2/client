package config

import (
	"net/http"

	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

func getConfigHandler(config *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.WriteResultToResponse(w, 200, config)
	}
}

func updateConfigHandler(config *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newUpdateConfig(r)
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

		updates := &types.Config{
			ChainID:         body.ChainID,
			RPCAddress:      body.RPCAddress,
			VerifierDir:     body.VerifierDir,
			TrustNode:       body.TrustNode,
			KeysDir:         body.KeysDir,
			ResolverAddress: body.ResolverAddress,
			KillSwitch:      body.KillSwitch,
		}

		if err := config.UpdateHook(updates); err != nil {
			utils.WriteErrorToResponse(w, 500, &types.Error{
				Message: "failed to call the config update hook",
				Info:    err.Error(),
			})
			return
		}

		config.Update(updates)
		if err := config.SaveToPath(""); err != nil {
			utils.WriteErrorToResponse(w, 500, &types.Error{
				Message: "failed to save the config",
				Info:    err.Error(),
			})
			return
		}

		utils.WriteResultToResponse(w, 200, config)
	}
}
