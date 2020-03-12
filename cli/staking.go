package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

func (cli *CLI) GetDelegatorValidatorsFromRPC(address sdk.AccAddress) (staking.Validators, error) {
	params := staking.NewQueryDelegatorParams(address)

	bz, err := cli.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorValidators), bz)
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

func (cli *CLI) GetAllValidators(r *http.Request) (interface{}, error) {
	status := r.URL.Query().Get("status")

	if cli.ExplorerAddress == "" {
		return nil, errors.New("invalid explorer address")
	}
	url := "http://" + cli.ExplorerAddress + "/validators"
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

func (cli *CLI) GetValidator(address string) (interface{}, error) {
	if cli.ExplorerAddress == "" {
		return nil, errors.New("invalid explorer address")
	}

	url := "http://" + cli.ExplorerAddress + "/validators/" + address

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
