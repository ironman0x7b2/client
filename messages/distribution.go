package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/tendermint/tendermint/libs/common"
)

type Rewards struct {
	FromAddress string `json:"from_address"`
	ValAddress  string `json:"val_address"`
}

func NewRewards(fromAddress, valAddress string) *Rewards {
	return &Rewards{
		FromAddress: fromAddress,
		ValAddress:  valAddress,
	}
}

func NewRewardsFromRaw(m *distribution.MsgWithdrawDelegatorReward) *Rewards {
	return &Rewards{
		FromAddress: common.HexBytes(m.DelegatorAddress.Bytes()).String(),
		ValAddress:  common.HexBytes(m.ValidatorAddress).String(),
	}
}

func (d *Rewards) Raw() (rewards distribution.MsgWithdrawDelegatorReward, err error) {
	rewards.DelegatorAddress, err = sdk.AccAddressFromHex(d.FromAddress)
	if err != nil {
		return rewards, err
	}

	rewards.ValidatorAddress, err = sdk.ValAddressFromHex(d.ValAddress)
	if err != nil {
		return rewards, err
	}

	return rewards, nil
}
