package types

import (
	"log"
	"os"
	"path/filepath"
)

// nolint: gochecknoglobals
var (
	DefaultConfigDir       string
	DefaultConfigFilePath  string
	DefaultChainID         = "sentinel-turing-2"
	DefaultRPCAddress      = "rpc.turing.sentinel.co:80"
	DefaultExplorerAddress = "145.239.224.179:8001"
)

// nolint: gochecknoinits
func init() {
	HomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultConfigDir = filepath.Join(HomeDir, ".sentinel", "client")
	DefaultConfigFilePath = filepath.Join(DefaultConfigDir, "config.json")
	
	if err := os.MkdirAll(DefaultConfigDir, os.ModePerm); err != nil {
		log.Panicln(err)
	}
}
