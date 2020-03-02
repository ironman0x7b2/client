package tx

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/utils"
)

/**
 * @api {get} /txs/ get all transactions
 * @apiDescription Used to get all transactions details
 * @apiName GetTransactions
 * @apiGroup tx
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getTxs(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		txs, err := cli.GetTxs(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)

			log.Println(err.Message)
			return
		}

		utils.WriteResultToResponse(w, 200, txs)
	}
}

/**
 * @api {get} /txs/bank/{address} get all bank transactions
 * @apiDescription Used to get all bank transactions details
 * @apiName GetBankTransactions
 * @apiGroup tx
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getBankTxs(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		txs, err := cli.GetBankTxs(vars["address"], r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)
			return
		}

		utils.WriteResultToResponse(w, 200, txs)
	}
}

/**
 * @api {get} /txs/{hash} get transaction details
 * @apiDescription Used to get transaction details
 * @apiName GetTransaction
 * @apiGroup tx
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getTx(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		tx, err := cli.GetTx(vars["hash"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)
			return
		}

		utils.WriteResultToResponse(w, 200, tx)
	}
}
