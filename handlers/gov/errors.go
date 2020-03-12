package gov

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	GOV = "gov"

	errFailedToGetProposal      = 12
	errFailedToGetProposals     = 13
	errFailedToGetProposalVotes = 14
	errFailedToGetProposalVote  = 15

	errMsgFailedToGetProposal      = "failed to get proposal"
	errMsgFailedToGetProposals     = "failed to get proposal"
	errMsgFailedToGetProposalVotes = "failed to get proposal votes"
	errMsgFailedToGetProposalVote  = "failed to get proposal vote"
)

func errorFailedToGetProposal() *types.Error {
	return types.NewError(GOV, errFailedToGetProposal, errMsgFailedToGetProposal)
}
func errorFailedToGetProposals() *types.Error {
	return types.NewError(GOV, errFailedToGetProposals, errMsgFailedToGetProposals)
}
func errorFailedToGetProposalVotes() *types.Error {
	return types.NewError(GOV, errFailedToGetProposalVotes, errMsgFailedToGetProposalVotes)
}
func errorFailedToGetProposalVote() *types.Error {
	return types.NewError(GOV, errFailedToGetProposalVote, errMsgFailedToGetProposalVote)
}
