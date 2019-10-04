package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/ironman0x7b2/client/types"
)

type MsgDelegate struct {
	DelegatorAddress string     `json:"delegator_address"`
	ValidatorAddress string     `json:"validator_address"`
	Amount           types.Coin `json:"amount"`
}

func NewMsgDelegate(delAddr string, valAddr string, amount types.Coin) *MsgDelegate {
	return &MsgDelegate{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
		Amount:           amount,
	}
}

func (d *MsgDelegate) Raw() (*staking.MsgDelegate, error) {
	var delegator staking.MsgDelegate
	var err error
	delegator.DelegatorAddress, err = sdk.AccAddressFromHex(d.DelegatorAddress)
	if err != nil {
		return &delegator, err
	}

	delegator.ValidatorAddress, err = sdk.ValAddressFromHex(d.ValidatorAddress)
	if err != nil {
		return &delegator, err
	}

	delegator.Amount = d.Amount.Raw()
	return &delegator, nil
}
