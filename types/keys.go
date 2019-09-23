package types

import (
	"log"
	"os"
	"path/filepath"
)

// nolint: gochecknoglobals
var (
	HomeDir               = os.ExpandEnv("$HOME")
	DefaultConfigDir      = filepath.Join(HomeDir, ".sentinel", "client")
	DefaultConfigFilePath = filepath.Join(DefaultConfigDir, "config.json")
)

// nolint: gochecknoinits
func init() {
	if err := os.MkdirAll(DefaultConfigDir, os.ModePerm); err != nil {
		log.Panicln(err)
	}
}
