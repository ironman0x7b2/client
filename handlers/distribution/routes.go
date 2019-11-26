package distribution

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("Withdraw-rewards").
		Methods("POST").Path("/accounts/{validatorAddress}/rewards").
		HandlerFunc(rewardsHandler(cli))
}
