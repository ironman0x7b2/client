package key

import (
	"net/http"

	"github.com/cosmos/go-bip39"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/models"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

func getKeysHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		infos, err := cli.Keybase.List()
		if err != nil {
			utils.WriteErrorToResponse(w, 500, &types.Error{
				Message: "failed to list the keys",
				Info:    err.Error(),
			})
			return
		}

		mnemonics := make([]string, len(infos))

		_keys := models.NewKeysFromRaw(infos, mnemonics)
		utils.WriteResultToResponse(w, 200, _keys)
	}
}

func addKeyHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newAddKey(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate the request body",
				Info:    err.Error(),
			})
			return
		}

		if _, err = cli.Keybase.Get(body.Name); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "duplicate key name",
				Info:    err.Error(),
			})
			return
		}

		if body.Mnemonic == "" {
			entropy, err := bip39.NewEntropy(256) // nolint: govet
			if err != nil {
				utils.WriteErrorToResponse(w, 400, &types.Error{
					Message: "failed to create the new entropy",
					Info:    err.Error(),
				})
				return
			}

			body.Mnemonic, err = bip39.NewMnemonic(entropy)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, &types.Error{
					Message: "failed to create the new mnemonic",
					Info:    err.Error(),
				})
				return
			}
		}

		info, err := cli.Keybase.CreateAccount(body.Name, body.Mnemonic, body.BIP39Password, body.Password, 0, 0)
		if err != nil {
			utils.WriteErrorToResponse(w, 500, &types.Error{
				Message: "failed to create the key",
				Info:    err.Error(),
			})
			return
		}

		key := models.NewKeyFromRaw(info, body.Mnemonic)
		utils.WriteResultToResponse(w, 201, key)
	}
}

func deleteKeyHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newDeleteKey(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate the request body",
				Info:    err.Error(),
			})
			return
		}

		if err = cli.Keybase.Delete(body.Name, body.Password, false); err != nil {
			utils.WriteErrorToResponse(w, 500, &types.Error{
				Message: "failed to delete the key",
				Info:    err.Error(),
			})
			return
		}

		utils.WriteResultToResponse(w, 200, nil)
	}
}
