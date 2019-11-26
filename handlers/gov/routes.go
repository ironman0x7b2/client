package gov

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("SubmitProposals").
		Methods("POST").Path("/proposals").
		HandlerFunc(submitProposalHandler(cli))
	r.Name("ProposalDeposit").
		Methods("POST").Path("/proposals/{id}/deposits").
		HandlerFunc(proposalDepositsHandler(cli))
	r.Name("SubmitProposals").
		Methods("POST").Path("/proposals").
		HandlerFunc(submitProposalHandler(cli))
}
