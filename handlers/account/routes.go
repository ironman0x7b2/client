package account

import (
	"github.com/gorilla/mux"
	
	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("GetAccount").
		Methods("GET").Path("/accounts/{address}").
		HandlerFunc(getAccountHandler(cli))
	r.Name("TransferCoins").
		Methods("POST").Path("/transfer").
		HandlerFunc(transferCoinsHandler(cli))
	
	r.Name("GetAllDelegatorDelegations").
		Methods("GET").Path("/accounts/{address}/delegations").
		HandlerFunc(getDelegatorDelegationsHandler(cli))
	r.Name("GetAllDelegatorValidators").
		Methods("GET").Path("/accounts/{address}/delegations/validators").
		HandlerFunc(getDelegatorValidatorsHandler(cli))
}
