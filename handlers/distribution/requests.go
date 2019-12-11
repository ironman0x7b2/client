package distribution

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ironman0x7b2/client/types"
)

type allRewards struct {
	From        string `json:"from"`
	FromAddress string `json:"from_address"`

	Memo          string         `json:"memo"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Gas           uint64         `json:"gas"`
	GasAdjustment float64        `json:"gas_adjustment"`

	Password string `json:"password"`
}

func newAllRewards(r *http.Request) (*rewards, error) {
	var body rewards
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (r *allRewards) Validate() error {
	if r.From == "" {
		return fmt.Errorf("invalid field from")
	}
	if r.FromAddress == "" {
		return fmt.Errorf("invalid field from_address")
	}

	if r.Password == "" {
		return fmt.Errorf("invalid field password")
	}

	return nil
}

type rewards struct {
	From        string `json:"from"`
	FromAddress string `json:"from_address"`

	Memo          string         `json:"memo"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Gas           uint64         `json:"gas"`
	GasAdjustment float64        `json:"gas_adjustment"`

	Password string `json:"password"`
}

func newRewards(r *http.Request) (*rewards, error) {
	var body rewards
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (r *rewards) Validate() error {
	if r.From == "" {
		return fmt.Errorf("invalid field from")
	}
	if r.FromAddress == "" {
		return fmt.Errorf("invalid field from_address")
	}

	if r.Password == "" {
		return fmt.Errorf("invalid field password")
	}

	return nil
}
