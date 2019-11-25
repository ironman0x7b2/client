package staking

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("GetAllDelegatorDelegations").
		Methods("GET").Path("/accounts/{address}/delegations").
		HandlerFunc(getDelegatorDelegations(cli))
	r.Name("GetAllDelegatorValidators").
		Methods("GET").Path("/accounts/{address}/delegations/validators").
		HandlerFunc(getDelegatorValidators(cli))

	r.Name("GetAllValidators").
		Methods("GET").Path("/validators").
		HandlerFunc(getAllValidators(cli))
	r.Name("GetValidator").
		Methods("GET").Path("/validators/{address}").
		HandlerFunc(getValidator(cli))

	r.Name("Delegate").
		Methods("POST").Path("/delegations/{valAddress}").
		HandlerFunc(delegateHandler(cli))
	r.Name("ReDelegate").
		Methods("POST").Path("/reDelegations/{valSrcAddress}").
		HandlerFunc(reDelegateHandler(cli))
}
