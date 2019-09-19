package hooks

import (
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/types"
)

func ConfigUpdateHook(cfg *types.Config, cli *_cli.CLI) func() error {
	return func() error {
		c := sdk.GetConfig()
		c.SetBech32PrefixForAccount(cfg.Bech32PrefixAccAddr, cfg.Bech32PrefixAccPub)
		c.SetBech32PrefixForValidator(cfg.Bech32PrefixValAddr, cfg.Bech32PrefixValPub)
		c.SetBech32PrefixForConsensusNode(cfg.Bech32PrefixConsAddr, cfg.Bech32PrefixConsPub)

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
