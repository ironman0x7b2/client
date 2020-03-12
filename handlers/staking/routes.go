package staking

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
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
		Methods("PUT").Path("/delegation/{valSrcAddress}").
		HandlerFunc(reDelegationHandler(cli))
	r.Name("Unbond").
		Methods("DELETE").Path("/delegation/{validatorAddress}").
		HandlerFunc(unbondHandler(cli))
}
