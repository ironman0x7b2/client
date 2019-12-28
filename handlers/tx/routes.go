package tx

import (
	"github.com/gorilla/mux"
	
	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("GetTransactionsDetails").
		Methods("GET").Path("/txs").
		HandlerFunc(getTxs(cli))
	r.Name("GetBankTransactionDetails").
		Methods("GET").Path("/txs/bank/{address}").
		HandlerFunc(getBankTxs(cli))
	r.Name("GetTransactionDetails").
		Methods("GET").Path("/txs/{hash}").
		HandlerFunc(getTx(cli))
}
