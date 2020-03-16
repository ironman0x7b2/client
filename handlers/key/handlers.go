package key

import (
	"log"
	"net/http"
	"strings"

	"github.com/cosmos/go-bip39"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/handlers/common"
	"github.com/ironman0x7b2/client/models"
	"github.com/ironman0x7b2/client/utils"
)

const MODULE = "key"

/**
 * @api {get} /keys get keys
 * @apiDescription Used to get keys details
 * @apiName GetKeys
 * @apiGroup keys
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getKeysHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		infos, err := cli.Keybase.List()
		if err != nil {
			utils.WriteErrorToResponse(w, 500, errorFailedToListKeys())

			log.Println(err.Error())
			return
		}

		mnemonics := make([]string, len(infos))

		_keys := models.NewKeysFromRaw(infos, mnemonics)
		utils.WriteResultToResponse(w, 200, _keys)
	}
}

/**
 * @api {get} /mnemonic get mnemonic
 * @apiDescription Used to get mnemonic
 * @apiName GetMnemonic
 * @apiGroup keys
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func generateMnemonicHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entropy, err := bip39.NewEntropy(256)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errorInvalidMnemonic())

			log.Println(err.Error())
			return
		}

		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errorFailedToCreateMnemonic())

			log.Println(err.Error())
			return
		}

		utils.WriteResultToResponse(w, 201, mnemonic)
	}
}

/**
 * @api {post} /keys add keys
 * @apiDescription Used to create keys
 * @apiName AddKeys
 * @apiGroup keys
 * @apiParamExample {json} Request-Example:
 * {
 *	"name":"Name",
 *	"password":"password",
 *	"mnumonic":"mnumonic"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func addKeyHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newAddKey(r)
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

		info, _ := cli.Keybase.Get(body.Name)
		if info != nil {
			utils.WriteErrorToResponse(w, 400, errorDuplicateKeyName())

			return
		}

		mnemonic := strings.Split(body.Mnemonic, " ")
		if len(mnemonic) != 24 {
			utils.WriteErrorToResponse(w, 400, errorInvalidMnemonicLength())

			log.Println("invalid mnemonic length")
			return
		}

		info, err = cli.Keybase.CreateAccount(body.Name, body.Mnemonic, body.BIP39Password, body.Password, 0, 0)
		if err != nil {
			utils.WriteErrorToResponse(w, 500, errorFailedToCreateKey())

			log.Println(err.Error())
			return
		}

		key := models.NewKeyFromRaw(info, body.Mnemonic)
		utils.WriteResultToResponse(w, 201, key)
	}
}

/**
 * @api {delete} /keys delete keys
 * @apiDescription Used to delete keys
 * @apiName DeleteKeys
 * @apiGroup keys
 * @apiParamExample {json} Request-Example:
 * {
 *	"name":"Name",
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func deleteKeyHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newDeleteKey(r)
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

		if err = cli.Keybase.Delete(body.Name, body.Password, false); err != nil {
			utils.WriteErrorToResponse(w, 500, errorFailedToDeleteKey())

			log.Println(err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, nil)
	}
}
