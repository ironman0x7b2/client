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
	if s.ResolverID == "" {
		return fmt.Errorf("invalid field resolver_id")
	}

	return nil
}

type submitTxHashToNode struct {
	NodeIP   string `json:"node_ip"`
	NodePort uint64 `json:"node_port"`
	TxHash   string `json:"tx_hash"`
}

func newsubmitTxHashToNode(r *http.Request) (*submitTxHashToNode, error) {
	var body submitTxHashToNode
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (s *submitTxHashToNode) Validate() error {
	if s.NodeIP == "" {
		return fmt.Errorf("invalid field node_ip")
	}
	if s.NodePort < 0 {
		return fmt.Errorf("invalid field node_port")
	}
	if s.TxHash == "" {
		return fmt.Errorf("invalid field tx_hash")
	}

	return nil
}

type getVPNKey struct {
	NodeIP   string `json:"node_ip"`
	NodePort uint64 `json:"node_port"`
}

func newGetPVNKey(r *http.Request) (*getVPNKey, error) {
	var body getVPNKey
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (g *getVPNKey) Validate() error {
	if g.NodeIP == "" {
		return fmt.Errorf("invalid field node_ip")
	}
	if g.NodePort < 0 {
		return fmt.Errorf("invalid field node_port")
	}

	return nil
}

type connectVPN struct {
	NodeIP      string `json:"node_ip"`
	NodePort    uint64 `json:"node_port"`
	AccountName string `json:"account_name"`
	Password    string `json:"password"`
	Key         string `json:"key"`
}

func newConnectVPN(r *http.Request) (*connectVPN, error) {
	var body connectVPN
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *connectVPN) Validate() error {
	if c.NodeIP == "" {
		return fmt.Errorf("invalid field node_ip")
	}
	if c.NodePort < 0 {
		return fmt.Errorf("invalid field node_port")
	}
	if c.AccountName == "" {
		return fmt.Errorf("invalid field account_name")
	}
	if c.Password == "" {
		return fmt.Errorf("invalid field password")
	}
	if c.Key == "" {
		return fmt.Errorf("invalid field key")
	}

	return nil
}
