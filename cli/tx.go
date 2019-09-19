package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
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
