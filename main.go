package main

import (
	"flag"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/gorilla/mux"
	"github.com/sentinel-official/hub/app"
	tm "github.com/tendermint/tendermint/types"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/handlers/account"
	"github.com/ironman0x7b2/client/handlers/config"
	"github.com/ironman0x7b2/client/handlers/key"
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
	cfg := types.NewDefaultConfig()
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

	cfg.SetUpdateHook(hooks.ConfigUpdateHook(cfg, cli))

	router := mux.NewRouter()
	router.Use(middlewares.AddHeaders)
	router.Use(middlewares.Log)

	config.RegisterRoutes(router, cfg)
	key.RegisterRoutes(router, cli)
	account.RegisterRoutes(router, cli)

	panic(http.ListenAndServe(address, router))
}
