package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sentinel-official/hub/app"
	hub "github.com/sentinel-official/hub/types"
	tm "github.com/tendermint/tendermint/types"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/handlers/account"
	"github.com/ironman0x7b2/client/handlers/config"
	"github.com/ironman0x7b2/client/handlers/distribution"
	"github.com/ironman0x7b2/client/handlers/gov"
	"github.com/ironman0x7b2/client/handlers/key"
	"github.com/ironman0x7b2/client/handlers/staking"
	"github.com/ironman0x7b2/client/handlers/tx"
	"github.com/ironman0x7b2/client/handlers/vpn"
	"github.com/ironman0x7b2/client/hooks"
	"github.com/ironman0x7b2/client/middlewares"
	"github.com/ironman0x7b2/client/types"
)

// nolint:gochecknoglobals
var (
	address string
)

// nolint:gochecknoinits
func init() {
	flag.StringVar(&address, "address", "0.0.0.0:8000", "server listen address")
	flag.Parse()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	cfg := types.NewDefaultConfig()

	conf := sdk.GetConfig()
	conf.SetBech32PrefixForAccount(hub.Bech32PrefixAccAddr, hub.Bech32PrefixAccPub)
	conf.SetBech32PrefixForValidator(hub.Bech32PrefixValAddr, hub.Bech32PrefixValPub)
	conf.SetBech32PrefixForConsensusNode(hub.Bech32PrefixConsAddr, hub.Bech32PrefixConsPub)
	conf.Seal()

	if err := cfg.LoadFromPath(""); err != nil {
		panic(err)
	}
	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	cdc := app.MakeCodec()
	tm.RegisterEventDatas(cdc)

	kb, err := keys.NewKeyBaseFromDir(cfg.KeysDir)
	if err != nil {
		panic(err)
	}

	cli := _cli.NewCLI(cdc, kb)

	cfg.SetUpdateHook(hooks.ConfigUpdateHook(cli))
	err = cfg.UpdateHook(cfg)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.Use(middlewares.AddHeaders)
	router.Use(middlewares.Log)

	config.RegisterRoutes(router, cfg)
	key.RegisterRoutes(router, cli)
	account.RegisterRoutes(router, cli)
	staking.RegisterRoutes(router, cli)
	gov.RegisterRoutes(router, cli)
	distribution.RegisterRoutes(router, cli)
	tx.RegisterRoutes(router, cli)
	vpn.RegisterRoutes(router, cli)

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(router)

	panic(http.ListenAndServe(address, handler))
}
	