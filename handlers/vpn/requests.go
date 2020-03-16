package vpn

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ironman0x7b2/client/types"
)

type subscription struct {
	From        string `json:"from"`
	FromAddress string `json:"from_address"`

	Amount     types.Coin `json:"amount"`
	NodeID     string     `json:"node_id"`
	ResolverID string     `json:"resolver_id"`

	NodeIP   string `json:"node_ip"`
	NodePort uint64 `json:"node_port"`

	Memo          string         `json:"memo"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Gas           uint64         `json:"gas"`
	GasAdjustment float64        `json:"gas_adjustment"`

	Password string `json:"password"`
}

func newSubscription(r *http.Request) (*subscription, error) {
	var body subscription
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (s *subscription) Validate() error {
	if s.From == "" {
		return fmt.Errorf("invalid field from")
	}
	if s.FromAddress == "" {
		return fmt.Errorf("invalid field from_address")
	}
	if s.Amount.Denom == "" || s.Amount.Value <= 0 {
		return fmt.Errorf("invalid field amount")
	}
	if s.Password == "" {
		return fmt.Errorf("invalid field password")
	}
	if s.NodeID == "" {
		return fmt.Errorf("invalid field node_id")
	}
	if s.NodeIP == "" {
		return fmt.Errorf("invalid field node_ip")
	}
	if s.NodePort < 0 {
		return fmt.Errorf("invalid field node_port")
	}
	if s.ResolverID == "" {
		return fmt.Errorf("invalid field resolver_id")
	}

	return nil
}

type NewVPN struct {
	From        string `json:"from"`
	FromAddress string `json:"from_address"`

	NodeID     string `json:"node_id"`
	ResolverID string `json:"resolver_id"`

	NodeIP   string `json:"node_ip"`
	NodePort uint64 `json:"node_port"`

	Memo          string         `json:"memo"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Gas           uint64         `json:"gas"`
	GasAdjustment float64        `json:"gas_adjustment"`

	Password string `json:"password"`
}

func newVPN(r *http.Request) (*NewVPN, error) {
	var body NewVPN
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (n *NewVPN) Validate() error {
	if n.From == "" {
		return fmt.Errorf("invalid field from")
	}
	if n.FromAddress == "" {
		return fmt.Errorf("invalid field from_address")
	}
	if n.Password == "" {
		return fmt.Errorf("invalid field password")
	}
	if n.NodeID == "" {
		return fmt.Errorf("invalid field node_id")
	}
	if n.NodeIP == "" {
		return fmt.Errorf("invalid field node_ip")
	}
	if n.NodePort < 0 {
		return fmt.Errorf("invalid field node_port")
	}
	if n.ResolverID == "" {
		return fmt.Errorf("invalid field resolver_id")
	}

	return nil
}
