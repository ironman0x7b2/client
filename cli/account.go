package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/staking"
	
	"github.com/ironman0x7b2/client/types"
)

func (c *CLI) GetAccount(address sdk.AccAddress) (auth.Account, error) {
	bytes, err := c.Codec.MarshalJSON(auth.NewQueryAccountParams(address))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	
	res, _, err := c.QueryWithData(fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount), bytes)
	if err != nil {
		if err.Error() == errors.New("no RPC client defined").Error() {
			log.Println(err.Error())
			return nil, err
		}
		return nil, nil
	}
	
	var account auth.Account
	if err := c.Codec.UnmarshalJSON(res, &account); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	
	return account, nil
}

func (cli *CLI) GetDelegatorDelegations(address sdk.AccAddress) (staking.Delegations, *types.Error) {
	params := staking.NewQueryDelegatorParams(address)
	
	bz, err := cli.Codec.MarshalJSON(params)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to marshal params",
			Info:    err.Error(),
		}
	}
	
	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorDelegations), bz)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to query delegations",
			Info:    err.Error(),
		}
	}
	
	var delegations staking.Delegations
	err = json.Unmarshal(res, &delegations)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to unmarshal delegations",
			Info:    err.Error(),
		}
	}
	
	return delegations, nil
}

func (cli *CLI) GetDelegatorValidators(address string) (interface{}, *types.Error) {
	url := "http://" + cli.ExplorerAddress + "/accounts/" + address + "/delegations/validators"
	
	res, err := http.Get(url)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to get validators",
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
