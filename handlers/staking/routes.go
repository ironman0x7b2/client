package staking

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("GetAllDelegatorDelegations").
		Methods("GET").Path("/accounts/{address}/delegations").
		HandlerFunc(getDelegatorDelegationsHandler(cli))
	r.Name("GetAllDelegatorValidators").
		Methods("GET").Path("/accounts/{address}/delegations/validators").
		HandlerFunc(getDelegatorValidatorsHandler(cli))

	r.Name("GetAllValidators").
		Methods("GET").Path("/validators").
		HandlerFunc(getAllValidatorsHandler(cli))
	r.Name("GetValidator").
		Methods("GET").Path("/validators/{address}").
		HandlerFunc(getValidatorHandler(cli))

	r.Name("Delegate").
		Methods("POST").Path("/delegations/{validatorAddress}").
		HandlerFunc(delegationHandler(cli))
	r.Name("ReDelegate").
		Methods("PUT").Path("/re-delegation/{valSrcAddress}").
		HandlerFunc(reDelegationHandler(cli))
	r.Name("UnDelegate").
		Methods("DELETE").Path("/un-delegation/{validatorAddress}").
		HandlerFunc(unDelegationHandler(cli))
}
