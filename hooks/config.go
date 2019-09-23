package hooks

import (
	"github.com/cosmos/cosmos-sdk/client/keys"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/types"
)

func ConfigUpdateHook(cfg *types.Config, cli *_cli.CLI) func() error {
	return func() error {
		kb, err := keys.NewKeyBaseFromDir(cfg.KeysDir)
		if err != nil {
			return err
		}

		client, verifier, err := _cli.CreateVerifier(types.DefaultConfigDir, cfg.ChainID, cfg.RPCAddress)
		if err != nil {
			return err
		}

		cli.Keybase = kb
		cli.Client = client
		cli.NodeURI = cfg.RPCAddress
		cli.Verifier = verifier
		cli.FromName = cfg.KeyName

		return nil
	}
}
