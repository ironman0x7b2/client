package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/tendermint/tendermint/libs/common"

	"github.com/ironman0x7b2/client/types"
)

type Proposal struct {
	FromAddress    string      `json:"from_address"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	Type           string      `json:"type"`
	InitialDeposit types.Coins `json:"initial_deposit"`
}

func NewProposal(fromAddress, title, description, _type string, deposit types.Coins) *Proposal {
	return &Proposal{
		FromAddress:    fromAddress,
		Title:          title,
		Description:    description,
		Type:           _type,
		InitialDeposit: deposit,
	}
}

func NewProposalFromRaw(m *gov.MsgSubmitProposal) *Proposal {
	return &Proposal{
		FromAddress:    common.HexBytes(m.Proposer.Bytes()).String(),
		Title:          m.Content.GetTitle(),
		Description:    m.Content.GetDescription(),
		Type:           m.Content.ProposalType(),
		InitialDeposit: types.NewCoinsFromRaw(m.InitialDeposit),
	}
}

func (p *Proposal) Raw() (proposal gov.MsgSubmitProposal, err error) {
	proposal.Proposer, err = sdk.AccAddressFromHex(p.FromAddress)
	if err != nil {
		return proposal, err
	}

	proposal.Content = gov.ContentFromProposalType(p.Title, p.Description, p.Type)
	proposal.InitialDeposit = p.InitialDeposit.Raw()

	return proposal, nil
}
