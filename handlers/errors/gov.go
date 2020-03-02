package errors

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	GOV = "gov"

	errFailedToGetProposal             = 8
	errFailedToUnmarshallProposal      = 8
	errFailedToGetProposalVotes        = 8
	errFailedToUnmarshallProposalVotes = 8
	errFailedToGetProposalVote         = 8
	errFailedToUnmarshallProposalVote  = 8

	errMsgFailedToGetProposal             = "failed to get proposal"
	errMsgFailedToUnmarshallProposal      = "failed to unmarshall proposal"
	errMsgFailedToGetProposalVotes        = "failed to get proposal votes"
	errMsgFailedToUnmarshallProposalVotes = "failed to unmarshall proposal votes"
	errMsgFailedToGetProposalVote         = "failed to get proposal vote"
	errMsgFailedToUnmarshallProposalVote  = "failed to unmarshall proposal vote"
)

func ErrorFailedToGetProposal() *types.Error {
	return types.NewError(GOV, errFailedToGetProposal, errMsgFailedToGetProposal)
}
func ErrorFailedToUnmarshalProposal() *types.Error {
	return types.NewError(GOV, errFailedToUnmarshallProposal, errMsgFailedToUnmarshallProposal)
}
func ErrorFailedToGetProposalVotes() *types.Error {
	return types.NewError(GOV, errFailedToGetProposalVotes, errMsgFailedToGetProposalVotes)
}
func ErrorFailedToUnmarshalProposalVotes() *types.Error {
	return types.NewError(GOV, errFailedToUnmarshallProposalVotes, errMsgFailedToUnmarshallProposalVotes)
}
func ErrorFailedToGetProposalVote() *types.Error {
	return types.NewError(GOV, errFailedToGetProposalVote, errMsgFailedToGetProposalVote)
}
func ErrorFailedToUnmarshalProposalVote() *types.Error {
	return types.NewError(GOV, errFailedToUnmarshallProposalVote, errMsgFailedToUnmarshallProposalVote)
}
