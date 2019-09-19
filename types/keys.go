package types

import (
	"os"
	"path/filepath"
)

var (
	HomeDir               = os.ExpandEnv("$HOME")
	DefaultConfigDir      = filepath.Join(HomeDir, ".sentinel", "client")
	DefaultConfigFilePath = filepath.Join(DefaultConfigDir, "config.json")
)

func init() {
	if err := os.MkdirAll(DefaultConfigDir, 0755); err != nil {
		panic(err)
	}
}
