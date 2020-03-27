package cli

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/ironman0x7b2/client/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

const MODULE = "txs"

func (c *CLI) Tx(messages []sdk.Msg, memo string, gas uint64, gasAdjustment float64,
	prices sdk.DecCoins, fees sdk.Coins, password string) (*sdk.TxResponse, error) {
	key, err := c.Keybase.Get(c.FromName)
	if err != nil {
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

	account, err := c.GetAccount(key.GetAddress())
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, err
	}

	txb := auth.NewTxBuilder(utils.GetTxEncoder(c.Codec),
		account.GetAccountNumber(), account.GetSequence(), gas, gasAdjustment,
		false, c.Verifier.ChainID(), memo, fees, prices).
		WithKeybase(c.Keybase)

	tx, err := txb.BuildAndSign(c.FromName, password, messages)
	if err != nil {
		return nil, err
	}

	node, err := c.GetNode()
	if err != nil {
		return nil, err
	}

	result, err := node.BroadcastTxSync(tx)
	if err != nil {
		return nil, err
	}

	res := sdk.NewResponseFormatBroadcastTx(result)

	return &res, nil
}

func (c *CLI) GetTx(hash string) (interface{}, error) {
	if c.ExplorerAddress == "" {
		return nil, errors.New("invalid explorer address")
	}

	url := "http://" + c.ExplorerAddress + "/txs/" + hash

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tx interface{}
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *CLI) GetTxs(r *http.Request) (interface{}, error) {
	signers := r.URL.Query().Get("signers")

	if c.ExplorerAddress == "" {
		return nil, errors.New("invalid explorer address")
	}

	url := "http://" + c.ExplorerAddress + "/txs"
	if signers != "" {
		url = url + "?signers=" + signers
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tx interface{}
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *CLI) GetBankTxs(address string, r *http.Request) (interface{}, error) {
	_type := r.URL.Query().Get("type")

	if c.ExplorerAddress == "" {
		return nil, errors.New("invalid explorer address")
	}

	url := "http://" + c.ExplorerAddress + "/txs/bank/" + address
	if _type == "send" {
		url = url + "?type=send"
	}
	if _type == "receive" {
		url = url + "?type=receive"
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tx interface{}
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
