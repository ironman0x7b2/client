package distribution

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("Withdraw-all-withdrawRewards").
		Methods("POST").Path("/accounts/withdraw-all-withdrawRewards").
		HandlerFunc(withdrawAllRewardsHandler(cli))
	r.Name("Withdraw-withdrawRewards").
		Methods("POST").Path("/accounts/withdraw-withdrawRewards/{validatorAddress}").
		HandlerFunc(withdrawRewardsHandler(cli))
}
