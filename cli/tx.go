package cli

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/ironman0x7b2/client/handlers/errors"
	"github.com/ironman0x7b2/client/types"
)

const MODULE = "txs"

func (c *CLI) Tx(messages []sdk.Msg, memo string, gas uint64, gasAdjustment float64,
	prices sdk.DecCoins, fees sdk.Coins, password string) (*sdk.TxResponse, *types.Error) {
	key, err := c.Keybase.Get(c.FromName)
	if err != nil {
		return nil, errors.ErrorFailedToGetKeyInfo()
	}

	account, err := c.GetAccount(key.GetAddress())
	if err != nil {
		return nil, errors.ErrorQueryAccount()
	}
	if account == nil {
		return nil, errors.ErrorAccountDoesNotExist()
	}

	txb := auth.NewTxBuilder(utils.GetTxEncoder(c.Codec),
		account.GetAccountNumber(), account.GetSequence(), gas, gasAdjustment,
		false, c.Verifier.ChainID(), memo, fees, prices).
		WithKeybase(c.Keybase)

	tx, err := txb.BuildAndSign(c.FromName, password, messages)
	if err != nil {
		return nil, errors.ErrorSignTransactions()
	}

	node, err := c.GetNode()
	if err != nil {
		return nil, errors.ErrorGetRPCNode()
	}

	result, err := node.BroadcastTxSync(tx)
	if err != nil {
		return nil, errors.ErrorFailedToBroadcastTransaction()
	}

	res := sdk.NewResponseFormatBroadcastTx(result)

	return &res, nil
}

func (cli *CLI) GetTx(hash string) (interface{}, *types.Error) {
	if cli.ExplorerAddress == "" {
		return nil, errors.ErrorInvalidExplorerAddress()
	}

	url := "http://" + cli.ExplorerAddress + "/txs/" + hash

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.ErrorFailedToGetTransaction()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.ErrorFailedToReadResponseBody(MODULE)
	}

	var tx interface{}
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return nil, errors.ErrorFailedToUnmarshalTransaction()
	}

	return tx, nil
}

func (cli *CLI) GetTxs(r *http.Request) (interface{}, *types.Error) {
	signers := r.URL.Query().Get("signers")

	if cli.ExplorerAddress == "" {
		return nil, errors.ErrorInvalidExplorerAddress()
	}

	url := "http://" + cli.ExplorerAddress + "/txs"
	if signers != "" {
		url = url + "?signers=" + signers
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.ErrorFailedToGetTransactions()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.ErrorFailedToReadResponseBody(MODULE)
	}

	var tx interface{}
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return nil, errors.ErrorFailedToUnmarshalTransactions()
	}

	return tx, nil
}

func (cli *CLI) GetBankTxs(address string, r *http.Request) (interface{}, *types.Error) {
	_type := r.URL.Query().Get("type")

	if cli.ExplorerAddress == "" {
		return nil, errors.ErrorInvalidExplorerAddress()
	}

	url := "http://" + cli.ExplorerAddress + "/txs/bank/" + address
	if _type == "send" {
		url = url + "?type=send"
	}
	if _type == "receive" {
		url = url + "?type=receive"
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.ErrorFailedToGetTransactions()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.ErrorFailedToReadResponseBody(MODULE)
	}

	var tx interface{}
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return nil, errors.ErrorFailedToUnmarshalTransactions()
	}

	return tx, nil
}
