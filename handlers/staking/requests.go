package staking

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ironman0x7b2/client/types"
)

type delegation struct {
	From        string     `json:"from"`
	FromAddress string     `json:"from_address"`
	Amount      types.Coin `json:"amount"`

	Memo          string         `json:"memo"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Gas           uint64         `json:"gas"`
	GasAdjustment float64        `json:"gas_adjustment"`

	Password string `json:"password"`
}

func newDelegate(r *http.Request) (*delegation, error) {
	var body delegation
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (d *delegation) Validate() error {
	if d.From == "" {
		return fmt.Errorf("invalid field from")
	}
	if d.FromAddress == "" {
		return fmt.Errorf("invalid field from_address")
	}
	if d.Amount.Denom == "" || d.Amount.Value <= 0 {
		return fmt.Errorf("invalid field amount")
	}
	if d.Password == "" {
		return fmt.Errorf("invalid field password")
	}

	return nil
}

type reDelegation struct {
	From           string     `json:"from"`
	FromAddress    string     `json:"from_address"`
	ValDestAddress string     `json:"val_dest_address"`
	Amount         types.Coin `json:"amount"`

	Memo          string         `json:"memo"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Gas           uint64         `json:"gas"`
	GasAdjustment float64        `json:"gas_adjustment"`

	Password string `json:"password"`
}

func newReDelegation(r *http.Request) (*reDelegation, error) {
	var body reDelegation
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (r *reDelegation) Validate() error {
	if r.From == "" {
		return fmt.Errorf("invalid field from")
	}
	if r.FromAddress == "" {
		return fmt.Errorf("invalid field from_address")
	}
	if r.ValDestAddress == "" {
		return fmt.Errorf("invalid field val_dest_address")
	}
	if r.Amount.Denom == "" || r.Amount.Value <= 0 {
		return fmt.Errorf("invalid field amount")
	}
	if r.Password == "" {
		return fmt.Errorf("invalid field password")
	}

	return nil
}
