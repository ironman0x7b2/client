package cli

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

func (cli *CLI) GetAllProposals(limit uint64) (*gov.Proposals, error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryProposalsParams(0, limit, nil, nil))
	if err != nil {
		return nil, err
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposals), bytes)
	if err != nil {
		return nil, err
	}

	var proposals gov.Proposals
	if err := cli.Codec.UnmarshalJSON(res, &proposals); err != nil {
		return nil, err
	}

	return &proposals, nil
}

func (cli *CLI) GetProposal(id uint64) (*gov.Proposal, error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryProposalParams(id))
	if err != nil {
		return nil, err
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposal), bytes)
	if err != nil {
		return nil, err
	}

	var proposal gov.Proposal
	if err := cli.Codec.UnmarshalJSON(res, &proposal); err != nil {
		return nil, err
	}

	return &proposal, nil
}

func (cli *CLI) GetProposalVotes(id uint64) (*gov.Votes, error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryProposalParams(id))
	if err != nil {
		return nil, err
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVotes), bytes)
	if err != nil {
		return nil, err
	}

	var votes gov.Votes
	if err := cli.Codec.UnmarshalJSON(res, &votes); err != nil {
		return nil, err
	}

	return &votes, nil
}

func (cli *CLI) GetProposalVoter(id uint64, voter sdk.AccAddress) (*gov.Votes, error) {
	bytes, err := cli.Codec.MarshalJSON(gov.NewQueryVoteParams(id, voter))
	if err != nil {
		return nil, err
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVote), bytes)
	if err != nil {
		return nil, err
	}

	var votes gov.Votes
	if err := cli.Codec.UnmarshalJSON(res, &votes); err != nil {
		return nil, err
	}

	return &votes, nil
}
