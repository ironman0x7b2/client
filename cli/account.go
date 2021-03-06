package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/ironman0x7b2/client/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func (c *CLI) GetAccount(address sdk.AccAddress) (auth.Account, error) {
	bytes, err := c.Codec.MarshalJSON(auth.NewQueryAccountParams(address))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

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

func (c *CLI) GetDelegatorDelegations(address string) (interface{}, error) {
	url := "http://" + c.ExplorerAddress + "/accounts/" + address + "/delegations"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var delegations interface{}
	err = json.Unmarshal(body, &delegations)
	if err != nil {
		return nil, err
	}

	return delegations, nil
}

func (c *CLI) GetDelegatorValidators(address string) (interface{}, error) {
	url := "http://" + c.ExplorerAddress + "/accounts/" + address + "/delegations/validators"

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
