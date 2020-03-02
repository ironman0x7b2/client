package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/ironman0x7b2/client/handlers/errors"
	"github.com/ironman0x7b2/client/types"
)

func (cli *CLI) GetDelegatorValidatorsFromRPC(address sdk.AccAddress, modeule string) (staking.Validators, *types.Error) {
	params := staking.NewQueryDelegatorParams(address)

	bz, err := cli.Codec.MarshalJSON(params)
	if err != nil {
		return nil, errors.ErrorFailedToMarshalParams(modeule)
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorValidators), bz)
	if err != nil {
		return nil, errors.ErrorFailedToQueryValidators(modeule)
	}

	var validators staking.Validators
	err = json.Unmarshal(res, &validators)
	if err != nil {
		return nil, errors.ErrorFailedToUnmarshallDelegatorValidators(modeule)
	}

	return validators, nil
}

func (cli *CLI) GetAllValidators(r *http.Request, module string) (interface{}, *types.Error) {
	status := r.URL.Query().Get("status")

	if cli.ExplorerAddress == "" {
		return nil, errors.ErrorInvalidExplorerAddress()
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
		return nil, errors.ErrorFailedToGetValidator()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.ErrorFailedToReadResponseBody(module)
	}

	var validators interface{}
	err = json.Unmarshal(body, &validators)
	if err != nil {
		return nil, errors.ErrorFailedToUnmarshallValidators()
	}

	return validators, nil
}

func (cli *CLI) GetValidator(address string, module string) (interface{}, *types.Error) {
	if cli.ExplorerAddress == "" {
		return nil, errors.ErrorInvalidExplorerAddress()
	}

	url := "http://" + cli.ExplorerAddress + "/validators/" + address

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.ErrorInvalidExplorerAddress()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.ErrorFailedToReadResponseBody(module)
	}

	var validator interface{}
	err = json.Unmarshal(body, &validator)
	if err != nil {
		return nil, errors.ErrorFailedToUnmarshallValidator()
	}

	return validator, nil
}
