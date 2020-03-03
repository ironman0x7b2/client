module github.com/ironman0x7b2/client

go 1.13

require (
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cosmos/cosmos-sdk v0.37.4
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/cosmos/ledger-cosmos-go v0.11.1 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/websocket v1.4.1
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/prometheus/procfs v0.0.7 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20190826022208-cac0b30c2563 // indirect
	github.com/rs/cors v1.7.0
	github.com/sentinel-official/dvpn-node v0.0.0-20190626093235-82a693e5e6e4
	github.com/sentinel-official/hub v0.2.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15 // indirect
	github.com/tendermint/tendermint v0.32.8
	golang.org/x/crypto v0.0.0-20191119213627-4f8c1d86b1ba // indirect
	golang.org/x/net v0.0.0-20191119073136-fc4aabc6c914 // indirect
	golang.org/x/sys v0.0.0-20191120155948-bd437916bb0e // indirect
	google.golang.org/genproto v0.0.0-20191115221424-83cc0476cb11 // indirect
)

replace github.com/sentinel-official/hub v0.2.0 => github.com/bitsndbyts/hub v0.2.1-0.20200214135426-a3563349e7d6

replace github.com/sentinel-official/dvpn-node => github.com/bitsndbyts/dvpn-node v0.0.0-20200217130525-2f0563c6a556
