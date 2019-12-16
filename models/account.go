package models

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/libs/common"

	"github.com/ironman0x7b2/client/types"
)

type Account struct {
	Address  string      `json:"address"`
	PubKey   string      `json:"pub_key"`
	Coins    types.Coins `json:"coins"`
	Sequence uint64      `json:"sequence"`
	Number   uint64      `json:"number"`
}

func NewAccountFromRaw(a auth.Account) (account Account) {
	if a.GetPubKey() == nil {
		account.PubKey = ""
	} else {
		account.PubKey = common.HexBytes(a.GetPubKey().Bytes()).String()
	}
	return Account{
		Address:  common.HexBytes(a.GetAddress().Bytes()).String(),
		Coins:    types.NewCoinsFromRaw(a.GetCoins()),
		Sequence: a.GetSequence(),
		Number:   a.GetAccountNumber(),
	}
}
