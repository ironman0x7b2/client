package validators

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("GetValidators").
		Methods("GET").Path("/validators").
		HandlerFunc(getValidatorsHandler(cli))
	r.Name("Delegate").
		Methods("POST").Path("/delegate").
		HandlerFunc(delegateHandler(cli))
}
