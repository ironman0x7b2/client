package cli

import (
	"fmt"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/ironman0x7b2/client/models"
)

func (c *CLI) GetValidators(r *http.Request) (_validators []models.Validator, err error) {
	_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
	if err != nil {
		return nil, err
	}

	params := staking.NewQueryValidatorsParams(page, limit, sdk.BondStatusBonded)
	bytes, err := c.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	route := fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryValidators)
	res, height, err := c.QueryWithData(route, bytes)
	if err != nil {
		return nil, err
	}

	var validators staking.Validators
	if err := c.Codec.UnmarshalJSON(res, &validators); err != nil {
		return nil, err
	}

	for _, val := range validators {
		_validators = append(_validators, models.NewValidatorFromRaw(val))
	}

	c.CLIContext = c.CLIContext.WithHeight(height)

	return _validators, nil
}

func (c *CLI) GetDelegatorValidators(delegator string, r *http.Request) (_validators []models.Validator, err error) {
	delegatorAddr, err := sdk.AccAddressFromHex(delegator)
	if err != nil {
		return nil, err
	}

	params := staking.NewQueryDelegatorParams(delegatorAddr)

	bytes, err := c.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	route := fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorValidators)
	res, height, err := c.QueryWithData(route, bytes)
	if err != nil {
		return nil, err
	}

	var validators staking.Validators
	if err := c.Codec.UnmarshalJSON(res, &validators); err != nil {
		return nil, err
	}

	for _, val := range validators {
		_validators = append(_validators, models.NewValidatorFromRaw(val))
	}

	c.CLIContext = c.WithHeight(height)

	return _validators, nil
}
