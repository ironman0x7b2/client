package cli

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/ironman0x7b2/client/types"
)

func (c *CLI) Tx(messages []sdk.Msg, memo string, gas uint64, gasAdjustment float64,
	prices sdk.DecCoins, fees sdk.Coins, password string) (*sdk.TxResponse, error) {

	key, err := c.Keybase.Get(c.FromName)
	if err != nil {
		return nil, err
	}

	account, err := c.GetAccount(key.GetAddress())
	if err != nil {
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

func (cli *CLI) GetTx(hash string) (interface{}, *types.Error) {
	if cli.ExplorerAddress == "" {
		return nil, &types.Error{
			Message: "no explorer address defined",
			Info:    "",
		}
	}

	url := "http://" + cli.ExplorerAddress + "/txs/" + hash

	res, err := http.Get(url)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to get transaction",
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

	var tx interface{}
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return nil, &types.Error{
			Message: "failed to unmarshal transaction",
			Info:    err.Error(),
		}
	}

	return tx, nil
}
