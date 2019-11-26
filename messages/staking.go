package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/tendermint/tendermint/libs/common"

	"github.com/ironman0x7b2/client/types"
)

type Delegate struct {
	FromAddress string     `json:"from_address"`
	ValAddress  string     `json:"val_address"`
	Amount      types.Coin `json:"amount"`
}

func NewDelegate(fromAddress, valAddress string, amount types.Coin) *Delegate {
	return &Delegate{
		FromAddress: fromAddress,
		ValAddress:  valAddress,
		Amount:      amount,
	}
}

func NewDelegateFromRaw(m *staking.MsgDelegate) *Delegate {
	return &Delegate{
		FromAddress: common.HexBytes(m.DelegatorAddress.Bytes()).String(),
		ValAddress:  common.HexBytes(m.ValidatorAddress).String(),
		Amount:      types.NewCoinFromRaw(m.Amount),
	}
}

func (d *Delegate) Raw() (delegate staking.MsgDelegate, err error) {
	delegate.DelegatorAddress, err = sdk.AccAddressFromHex(d.FromAddress)
	if err != nil {
		return delegate, err
	}

	delegate.ValidatorAddress, err = sdk.ValAddressFromHex(d.ValAddress)
	if err != nil {
		return delegate, err
	}

	delegate.Amount = d.Amount.Raw()

	return delegate, nil
}

type ReDelegate struct {
	FromAddress    string     `json:"from_address"`
	ValSrcAddress  string     `json:"val_src_address"`
	ValDestAddress string     `json:"val_dest_address"`
	Amount         types.Coin `json:"amount"`
}

func NewReDelegate(fromAddress, valSrcAddress, valDestAddress string, amount types.Coin) *ReDelegate {
	return &ReDelegate{
		FromAddress:    fromAddress,
		ValSrcAddress:  valSrcAddress,
		ValDestAddress: valDestAddress,
		Amount:         amount,
	}
}

func NewReDelegateFromRaw(m *staking.MsgBeginRedelegate) *ReDelegate {
	return &ReDelegate{
		FromAddress:    common.HexBytes(m.DelegatorAddress.Bytes()).String(),
		ValSrcAddress:  common.HexBytes(m.ValidatorSrcAddress).String(),
		ValDestAddress: common.HexBytes(m.ValidatorDstAddress).String(),
		Amount:         types.NewCoinFromRaw(m.Amount),
	}
}

func (d *ReDelegate) Raw() (reDelegate staking.MsgBeginRedelegate, err error) {
	reDelegate.DelegatorAddress, err = sdk.AccAddressFromHex(d.FromAddress)
	if err != nil {
		return reDelegate, err
	}

	reDelegate.ValidatorSrcAddress, err = sdk.ValAddressFromHex(d.ValSrcAddress)
	if err != nil {
		return reDelegate, err
	}

	reDelegate.ValidatorDstAddress, err = sdk.ValAddressFromHex(d.ValDestAddress)
	if err != nil {
		return reDelegate, err
	}

	reDelegate.Amount = d.Amount.Raw()

	return reDelegate, nil
}

type UnDelegate struct {
	FromAddress string     `json:"from_address"`
	ValAddress  string     `json:"val_address"`
	Amount      types.Coin `json:"amount"`
}

func NewUnDelegate(fromAddress, valAddress string, amount types.Coin) *UnDelegate {
	return &UnDelegate{
		FromAddress: fromAddress,
		ValAddress:  valAddress,
		Amount:      amount,
	}
}

func NewUnDelegateFromRaw(m *staking.MsgUndelegate) *UnDelegate {
	return &UnDelegate{
		FromAddress: common.HexBytes(m.DelegatorAddress.Bytes()).String(),
		ValAddress:  common.HexBytes(m.ValidatorAddress).String(),
		Amount:      types.NewCoinFromRaw(m.Amount),
	}
}

func (d *UnDelegate) Raw() (undelegate staking.MsgUndelegate, err error) {
	undelegate.DelegatorAddress, err = sdk.AccAddressFromHex(d.FromAddress)
	if err != nil {
		return undelegate, err
	}

	undelegate.ValidatorAddress, err = sdk.ValAddressFromHex(d.ValAddress)
	if err != nil {
		return undelegate, err
	}

	undelegate.Amount = d.Amount.Raw()

	return undelegate, nil
}
