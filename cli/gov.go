package cli

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"

	"github.com/ironman0x7b2/client/handlers/errors"
	"github.com/ironman0x7b2/client/types"
)

func (cli *CLI) GetAllProposals(limit uint64, module string) (*gov.Proposals, *types.Error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryProposalsParams(0, limit, nil, nil))
	if err != nil {
		return nil, errors.ErrorFailedToMarshalParams(module)
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposals), bytes)
	if err != nil {
		return nil, errors.ErrorFailedToGetProposal()
	}

	var proposals gov.Proposals
	if err := cli.Codec.UnmarshalJSON(res, &proposals); err != nil {
		return nil, errors.ErrorFailedToUnmarshalProposal()
	}

	return &proposals, nil
}

func (cli *CLI) GetProposal(id uint64, module string) (*gov.Proposal, *types.Error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryProposalParams(id))
	if err != nil {
		return nil, errors.ErrorParseQueryParams(module)
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposal), bytes)
	if err != nil {
		return nil, errors.ErrorFailedToGetProposal()
	}

	var proposal gov.Proposal
	if err := cli.Codec.UnmarshalJSON(res, &proposal); err != nil {
		return nil, errors.ErrorFailedToUnmarshalProposal()
	}

	return &proposal, nil
}

func (cli *CLI) GetProposalVotes(id uint64, module string) (gov.Votes, *types.Error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryProposalParams(id))
	if err != nil {
		return nil, errors.ErrorParseQueryParams(module)
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVotes), bytes)
	if err != nil {
		return nil, errors.ErrorFailedToGetProposalVotes()
	}

	var votes gov.Votes
	if err := cli.Codec.UnmarshalJSON(res, &votes); err != nil {
		return nil, errors.ErrorFailedToUnmarshalProposalVotes()
	}

	return votes, nil
}

func (cli *CLI) GetProposalVote(id uint64, voter sdk.AccAddress) (gov.Vote, *types.Error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryVoteParams(id, voter))
	if err != nil {
		return gov.Vote{}, errors.ErrorParseQueryParams(MODULE)
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVote), bytes)
	if err != nil {
		return gov.Vote{}, errors.ErrorFailedToGetProposalVote()
	}

	var _vote gov.Vote
	if err := cli.Codec.UnmarshalJSON(res, &_vote); err != nil {
		return gov.Vote{}, errors.ErrorFailedToUnmarshalProposalVote()
	}

	return _vote, nil
}
