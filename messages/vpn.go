package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn"
	"github.com/tendermint/tendermint/libs/common"

	"github.com/ironman0x7b2/client/types"
)

type Subscription struct {
	FromAddress string     `json:"from_address"`
	Amount      types.Coin `json:"amount"`
	NodeID      string     `json:"node_id"`
	ResolverID  string     `json:"resolver_id"`
}

func NewSubscription(fromAddress string, amount types.Coin, nodeID, resolverID string) *Subscription {
	return &Subscription{
		FromAddress: fromAddress,
		Amount:      amount,
		NodeID:      nodeID,
		ResolverID:  resolverID,
	}
}

func NewSubscriptionFromRaw(m *vpn.MsgStartSubscription) *Subscription {
	return &Subscription{
		FromAddress: common.HexBytes(m.From.Bytes()).String(),
		Amount:      types.NewCoinFromRaw(m.Deposit),
		NodeID:      m.NodeID.String(),
		ResolverID:  m.ResolverID.String(),
	}
}

func (s *Subscription) Raw() (subscription vpn.MsgStartSubscription, err error) {
	subscription.From, err = sdk.AccAddressFromHex(s.FromAddress)
	if err != nil {
		return subscription, err
	}

	subscription.Deposit = s.Amount.Raw()
	id, _ := hub.NewNodeIDFromString(s.NodeID)
	_id, _ := hub.NewResolverIDFromString(s.ResolverID)
	subscription.NodeID = id
	subscription.ResolverID = _id

	return subscription, nil
}
