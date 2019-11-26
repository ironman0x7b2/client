package gov

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("SubmitProposals").
		Methods("POST").Path("/proposals").
		HandlerFunc(submitProposal(cli))
}
