package gov

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("GetAllProposals").
		Methods("GET").Path("/proposals").
		HandlerFunc(getAllProposalsHandler(cli))
	r.Name("GetProposal").
		Methods("GET").Path("/proposals/{id}").
		HandlerFunc(getProposalHandler(cli))
	r.Name("GetProposalVotes").
		Methods("GET").Path("/proposals/{id}/votes").
		HandlerFunc(getProposalVotesHandler(cli))
	r.Name("GetProposalVote").
		Methods("GET").Path("/proposals/{id}/voters/{address}").
		HandlerFunc(getProposalVoteHandler(cli))

	r.Name("SubmitProposals").
		Methods("POST").Path("/proposals").
		HandlerFunc(submitProposalHandler(cli))
	r.Name("ProposalDeposit").
		Methods("POST").Path("/proposals/{id}/deposits").
		HandlerFunc(proposalDepositsHandler(cli))
	r.Name("ProposalsVotes").
		Methods("POST").Path("/proposals/{id}/votes").
		HandlerFunc(proposalVotesHandler(cli))
}
