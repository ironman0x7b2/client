package cli

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"

	"github.com/ironman0x7b2/client/types"
)

func (cli *CLI) GetAllProposals(limit uint64) (*gov.Proposals, *types.Error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryProposalsParams(0, limit, nil, nil))
	if err != nil {
		return nil, &types.Error{
			Message: "failed to marshal query params",
			Info:    err.Error(),
		}
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposals), bytes)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to get proposals",
			Info:    err.Error(),
		}
	}

	var proposals gov.Proposals
	if err := cli.Codec.UnmarshalJSON(res, &proposals); err != nil {
		return nil, &types.Error{
			Message: "failed to unmarshal proposals",
			Info:    err.Error(),
		}
	}

	return &proposals, nil
}

func (cli *CLI) GetProposal(id uint64) (*gov.Proposal, *types.Error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryProposalParams(id))
	if err != nil {
		return nil, &types.Error{
			Message: "failed to marshal query params",
			Info:    err.Error(),
		}
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposal), bytes)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to get proposal",
			Info:    err.Error(),
		}
	}

	var proposal gov.Proposal
	if err := cli.Codec.UnmarshalJSON(res, &proposal); err != nil {
		return nil, &types.Error{
			Message: "failed to unmarshal proposal",
			Info:    err.Error(),
		}
	}

	return &proposal, nil
}

func (cli *CLI) GetProposalVotes(id uint64) (gov.Votes, *types.Error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryProposalParams(id))
	if err != nil {
		return nil, &types.Error{
			Message: "failed to marshal query params",
			Info:    err.Error(),
		}
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVotes), bytes)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to get proposal votes",
			Info:    err.Error(),
		}
	}

	var votes gov.Votes
	if err := cli.Codec.UnmarshalJSON(res, &votes); err != nil {
		return nil, &types.Error{
			Message: "failed to unmarshal votes",
			Info:    err.Error(),
		}
	}

	return votes, nil
}

func (cli *CLI) GetProposalVote(id uint64, voter sdk.AccAddress) (gov.Vote, *types.Error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryVoteParams(id, voter))
	if err != nil {
		return gov.Vote{}, &types.Error{
			Message: "failed to unmarshal query params",
			Info:    err.Error(),
		}
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVote), bytes)
	if err != nil {
		return gov.Vote{}, &types.Error{
			Message: "failed to get proposal voters",
			Info:    err.Error(),
		}
	}

	var _vote gov.Vote
	if err := cli.Codec.UnmarshalJSON(res, &_vote); err != nil {
		return gov.Vote{}, &types.Error{
			Message: "failed to unmarshal vote",
			Info:    err.Error(),
		}
	}

	return _vote, nil
}
