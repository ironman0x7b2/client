package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/ironman0x7b2/client/types"
)

func (cli *CLI) GetDelegatorValidatorsFromRPC(address sdk.AccAddress) (staking.Validators, *types.Error) {
	params := staking.NewQueryDelegatorParams(address)

	bz, err := cli.Codec.MarshalJSON(params)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to marshal params",
			Info:    err.Error(),
		}
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorValidators), bz)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to query delegator validators",
			Info:    err.Error(),
		}
	}

	var validators staking.Validators
	err = json.Unmarshal(res, &validators)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to unmarshal delegator validators",
			Info:    err.Error(),
		}
	}

	return validators, nil
}

func (cli *CLI) GetAllValidators(r *http.Request) (interface{}, *types.Error) {
	status := r.URL.Query().Get("status")

	if cli.ExplorerAddress == "" {
		return nil, &types.Error{
			Message: "no explorer address defined",
			Info:    "",
		}
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
		return nil, &types.Error{
			Message: "failed to get validator",
			Info:    err.Error(),
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to read response body",
			Info:    err.Error(),
		}
	}

	var validators interface{}
	err = json.Unmarshal(body, &validators)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to unmarshal validators",
			Info:    err.Error(),
		}
	}

	return validators, nil
}

func (cli *CLI) GetValidator(address string) (interface{}, *types.Error) {
	if cli.ExplorerAddress == "" {
		return nil, &types.Error{
			Message: "no explorer address defined",
			Info:    "",
		}
	}

	url := "http://" + cli.ExplorerAddress + "/validators/" + address

	res, err := http.Get(url)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to get validator",
			Info:    err.Error(),
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to read response body",
			Info:    err.Error(),
		}
	}

	var validator interface{}
	err = json.Unmarshal(body, &validator)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to unmarshal validator",
			Info:    err.Error(),
		}
	}

	return validator, nil
}
