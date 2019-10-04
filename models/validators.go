package models

import (
	"time"

	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/tendermint/tendermint/libs/common"

	"github.com/ironman0x7b2/client/types"
)

type Validator struct {
	Address          string `json:"address"`
	PubKey           string `json:"pub_key"`
	ConsensusAddress string `json:"consensus_address"`
	ConsensusPubKey  string `json:"consensus_pub_key"`

	Description types.ValidatorDescription `json:"description"`
	Commission  types.ValidatorCommission  `json:"commission"`

	Jailed     bool   `json:"jailed"`
	BondStatus string `json:"bond_status"`

	Amount            types.Coin `json:"amount"`
	DelegatorShares   string     `json:"delegator_shares"`
	MinSelfDelegation int64      `json:"min_self_delegation"`

	UnbondingHeight         int64     `json:"unbonding_height"`
	UnbondingCompletionTime time.Time `json:"unbonding_completion_time"`

	Power    int64 `json:"power"`
	Priority int64 `json:"priority"`
}

func NewValidatorFromRaw(v staking.Validator) Validator {
	return Validator{
		Address:          common.HexBytes(v.OperatorAddress.Bytes()).String(),
		PubKey:           "",
		ConsensusAddress: common.HexBytes(v.ConsPubKey.Address().Bytes()).String(),
		ConsensusPubKey:  common.HexBytes(v.ConsPubKey.Bytes()).String(),
		Description: types.ValidatorDescription{
			Moniker:  v.Description.Moniker,
			Identity: v.Description.Identity,
			Website:  v.Description.Website,
			Details:  v.Description.Details,
		},
		Commission: types.ValidatorCommission{
			Rate:          v.Commission.Rate.String(),
			MaxRate:       v.Commission.MaxRate.String(),
			MaxChangeRate: v.Commission.MaxChangeRate.String(),
			UpdatedAt:     v.Commission.UpdateTime,
		},
		Jailed:     v.Jailed,
		BondStatus: v.Status.String(),
		Amount: types.Coin{
			Denom: "",
			Value: v.Tokens.Int64(),
		},
		DelegatorShares:         v.DelegatorShares.String(),
		MinSelfDelegation:       v.MinSelfDelegation.Int64(),
		UnbondingHeight:         v.UnbondingHeight,
		UnbondingCompletionTime: v.UnbondingCompletionTime,
		Power:                   0,
		Priority:                0,
	}
}
