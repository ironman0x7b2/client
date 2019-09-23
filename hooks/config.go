package hooks

import (
	"github.com/cosmos/cosmos-sdk/client/keys"
	kb "github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/tendermint/tendermint/lite"
	"github.com/tendermint/tendermint/rpc/client"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/types"
)

func ConfigUpdateHook(cli *_cli.CLI) func(*types.Config) error {
	return func(c *types.Config) (err error) {
		var _kb kb.Keybase

		if c.KeysDir != "" {
			_kb, err = keys.NewKeyBaseFromDir(c.KeysDir)
			if err != nil {
				return err
			}
		}

		var _client *client.HTTP
		var _verifier *lite.DynamicVerifier

		if c.VerifierDir != "" || c.ChainID != "" || c.RPCAddress != "" {
			_client, _verifier, err = _cli.NewVerifier(c.VerifierDir, c.ChainID, c.RPCAddress)
			if err != nil {
				return err
			}
		}

		if c.KeysDir != "" {
			cli.Keybase = _kb
		}

		if c.VerifierDir != "" || c.ChainID != "" || c.RPCAddress != "" {
			cli.Client = _client
			cli.Verifier = _verifier
			cli.NodeURI = c.RPCAddress
			cli.VerifierHome = c.VerifierDir
		}

		cli.TrustNode = c.TrustNode

		return nil
	}
}
