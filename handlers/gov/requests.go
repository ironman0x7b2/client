package gov

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ironman0x7b2/client/types"
)

type proposal struct {
	From        string      `json:"from"`
	FromAddress string      `json:"from_address"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Amount      types.Coins `json:"amount"`

	Memo          string         `json:"memo"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Gas           uint64         `json:"gas"`
	GasAdjustment float64        `json:"gas_adjustment"`

	Password string `json:"password"`
}

func newProposal(r *http.Request) (*proposal, error) {
	var body proposal
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (p *proposal) Validate() error {
	if p.From == "" {
		return fmt.Errorf("invalid field from")
	}
	if p.FromAddress == "" {
		return fmt.Errorf("invalid field from_address")
	}
	if p.Title == "" {
		return fmt.Errorf("invalid field title")
	}
	if p.Description == "" {
		return fmt.Errorf("invalid field description")
	}
	if p.Type == "" {
		return fmt.Errorf("invalid field type")
	}
	if len(p.Amount) == 0 {
		return fmt.Errorf("invalid field amount")
	}
	if p.Password == "" {
		return fmt.Errorf("invalid field password")
	}

	return nil
}
