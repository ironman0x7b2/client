package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/ironman0x7b2/client/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

func (c *CLI) GetDelegatorValidatorsFromRPC(address sdk.AccAddress) (staking.Validators, error) {
	if reflect.ValueOf(c.Verifier).IsNil() {
		cfg := types.NewDefaultConfig()
		client, verifier, err := NewVerifier(cfg.VerifierDir, cfg.ChainID, cfg.RPCAddress)
		if err != nil {
			return nil, err
		}
		if reflect.ValueOf(verifier).IsNil() {
			return nil, errors.New("Error while connectiong rpc")
		} else {
			c.Client = client
			c.Verifier = verifier
		}
	}

	params := staking.NewQueryDelegatorParams(address)

	bz, err := c.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, _, err := c.QueryWithData(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorValidators), bz)
	if err != nil {
		return nil, err
	}

	var validators staking.Validators
	err = json.Unmarshal(res, &validators)
	if err != nil {
		return nil, err
	}

	return validators, nil
}

func (c *CLI) GetAllValidators(r *http.Request) (interface{}, error) {
	status := r.URL.Query().Get("status")

	if c.ExplorerAddress == "" {
		return nil, errors.New("invalid explorer address")
	}
	url := "http://" + c.ExplorerAddress + "/validators"
	if status == "active" {
		url = url + "?status=active"
	}
	if status == "inactive" {
		url = url + "?status=inactive"
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var validators interface{}
	err = json.Unmarshal(body, &validators)
	if err != nil {
		return nil, err
	}

	return validators, nil
}

func (c *CLI) GetValidator(address string) (interface{}, error) {
	if c.ExplorerAddress == "" {
		return nil, errors.New("invalid explorer address")
	}

	url := "http://" + c.ExplorerAddress + "/validators/" + address

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("no RPC client defined")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var validator interface{}
	err = json.Unmarshal(body, &validator)
	if err != nil {
		return nil, err
	}

	return validator, nil
}
