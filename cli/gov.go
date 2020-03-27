package cli

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ironman0x7b2/client/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

func (c *CLI) GetAllProposals(limit uint64) (*gov.Proposals, error) {
	if reflect.ValueOf(c.Verifier).IsNil() {
		cfg := types.NewDefaultConfig()
		client, verifier, err := NewVerifier(cfg.VerifierDir, cfg.ChainID, cfg.RPCAddress)
		if err != nil {
			return nil, err
		}

		if reflect.ValueOf(verifier).IsNil() {
			return nil, errors.New("Error while connectiong rpc")
		} else {
			c.Client = client
			c.Verifier = verifier
		}
	}

	bytes, err := c.Codec.MarshalJSON(gov.NewQueryProposalsParams(0, limit, nil, nil))
	if err != nil {
		return nil, err
	}

	res, _, err := c.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposals), bytes)
	if err != nil {
		return nil, err
	}

	var proposals gov.Proposals
	if err := c.Codec.UnmarshalJSON(res, &proposals); err != nil {
		return nil, err
	}

	return &proposals, nil
}

func (c *CLI) GetProposal(id uint64) (*gov.Proposal, error) {
	if reflect.ValueOf(c.Verifier).IsNil() {
		cfg := types.NewDefaultConfig()
		client, verifier, err := NewVerifier(cfg.VerifierDir, cfg.ChainID, cfg.RPCAddress)
		if err != nil {
			return nil, err
		}
		if reflect.ValueOf(verifier).IsNil() {
			return nil, errors.New("Error while connectiong rpc")
		} else {
			c.Client = client
			c.Verifier = verifier
		}
	}

	bytes, err := c.Codec.MarshalJSON(gov.NewQueryProposalParams(id))
	if err != nil {
		return nil, err
	}

	res, _, err := c.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposal), bytes)
	if err != nil {
		return nil, err
	}

	var proposal gov.Proposal
	if err := c.Codec.UnmarshalJSON(res, &proposal); err != nil {
		return nil, err
	}

	return &proposal, nil
}

func (c *CLI) GetProposalVotes(id uint64) (gov.Votes, error) {
	if reflect.ValueOf(c.Verifier).IsNil() {
		cfg := types.NewDefaultConfig()
		client, verifier, err := NewVerifier(cfg.VerifierDir, cfg.ChainID, cfg.RPCAddress)
		if err != nil {
			return nil, err
		}
		if reflect.ValueOf(verifier).IsNil() {
			return nil, errors.New("Error while connectiong rpc")
		} else {
			c.Client = client
			c.Verifier = verifier
		}
	}

	bytes, err := c.Codec.MarshalJSON(gov.NewQueryProposalParams(id))
	if err != nil {
		return nil, err
	}

	res, _, err := c.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVotes), bytes)
	if err != nil {
		return nil, err
	}

	var votes gov.Votes
	if err := c.Codec.UnmarshalJSON(res, &votes); err != nil {
		return nil, err
	}

	return votes, nil
}

func (c *CLI) GetProposalVote(id uint64, voter sdk.AccAddress) (gov.Vote, error) {
	if reflect.ValueOf(c.Verifier).IsNil() {
		cfg := types.NewDefaultConfig()
		client, verifier, err := NewVerifier(cfg.VerifierDir, cfg.ChainID, cfg.RPCAddress)
		if err != nil {
			return gov.Vote{}, err
		}
		if reflect.ValueOf(verifier).IsNil() {
			return gov.Vote{}, errors.New("Error while connectiong rpc")
		} else {
			c.Client = client
			c.Verifier = verifier
		}
	}

	bytes, err := c.Codec.MarshalJSON(gov.NewQueryVoteParams(id, voter))
	if err != nil {
		return gov.Vote{}, err
	}

	res, _, err := c.QueryWithData(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVote), bytes)
	if err != nil {
		return gov.Vote{}, err
	}

	var _vote gov.Vote
	if err := c.Codec.UnmarshalJSON(res, &_vote); err != nil {
		return gov.Vote{}, err
	}

	return _vote, nil
}
