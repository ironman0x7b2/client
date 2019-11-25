package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ironman0x7b2/client/types"
)

type transferCoins struct {
	From        string      `json:"from"`
	FromAddress string      `json:"from_address"`
	ToAddress   string      `json:"to_address"`
	Amount      types.Coins `json:"amount"`

	Memo          string         `json:"memo"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Gas           uint64         `json:"gas"`
	GasAdjustment float64        `json:"gas_adjustment"`

	Password string `json:"password"`
}

func newTransferCoins(r *http.Request) (*transferCoins, error) {
	var body transferCoins
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (t *transferCoins) Validate() error {
	if t.From == "" {
		return fmt.Errorf("invalid field from")
	}
	if t.FromAddress == "" {
		return fmt.Errorf("invalid field from_address")
	}
	if t.ToAddress == "" {
		return fmt.Errorf("invalid field to_address")
	}
	if len(t.Amount) == 0 {
		return fmt.Errorf("invalid field amount")
	}
	if t.Password == "" {
		return fmt.Errorf("invalid field password")
	}

	return nil
}
