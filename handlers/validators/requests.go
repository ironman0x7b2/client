package validators

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ironman0x7b2/client/types"
)

type delegateCoins struct {
	From             string     `json:"from"`
	FromAddress      string     `json:"from_address"`
	DelegatorAddress string     `json:"delegator_address"`
	ValidatorAddress string     `json:"validator_address"`
	Amount           types.Coin `json:"amount"`

	Memo          string         `json:"memo"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Gas           uint64         `json:"gas"`
	GasAdjustment float64        `json:"gas_adjustment"`

	Password string `json:"password"`
}

func newDelegateCoins(r *http.Request) (*delegateCoins, error) {
	var body delegateCoins
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (d *delegateCoins) Validate() error {
	if d.FromAddress == "" {
		return fmt.Errorf("invalid field from_address")
	}
	if d.DelegatorAddress == "" {
		return fmt.Errorf("invalid field delegator_address")
	}
	if d.ValidatorAddress == "" {
		return fmt.Errorf("invalid field validator_address")
	}
	if d.Password == "" {
		return fmt.Errorf("invalid field password")
	}

	return nil
}
