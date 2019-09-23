package types

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sentinel-official/hub/app"
)

type Config struct {
	ChainID    string `json:"chain_id"`
	RPCAddress string `json:"rpc_address"`
	KeysDir    string `json:"keys_dir"`
	KeyName    string `json:"key_name"`

	uh func() error
}

func NewDefaultConfig() *Config {
	return &Config{
		ChainID:    "sentinel-turing-1",
		RPCAddress: "127.0.0.1:26657",
		KeysDir:    app.DefaultCLIHome,
		KeyName:    "",
	}
}

func (c *Config) Update(cfg *Config) {
	if cfg.ChainID != "" {
		c.ChainID = cfg.ChainID
	}
	if cfg.RPCAddress != "" {
		c.RPCAddress = cfg.RPCAddress
	}
	if cfg.KeysDir != "" {
		c.KeysDir = cfg.KeysDir
	}
	if cfg.KeyName != "" {
		c.KeyName = cfg.KeyName
	}
}

func (c Config) Copy() Config {
	return c
}

func (c *Config) SetUpdateHook(hook func() error) {
	c.uh = hook
}

func (c *Config) UpdateHook() error {
	if err := c.uh(); err != nil {
		return err
	}

	return c.SaveToPath("")
}

func (c *Config) LoadFromPath(path string) error {
	if path == "" {
		path = DefaultConfigFilePath
	}

	if _, err := os.Stat(path); err != nil {
		if err := c.SaveToPath(path); err != nil {
			return err
		}
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*c = Config{}
		return nil
	}

	return json.Unmarshal(data, c)
}

func (c *Config) SaveToPath(path string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	if path == "" {
		path = DefaultConfigFilePath
	}

	return ioutil.WriteFile(path, bytes, os.ModePerm)
}

func (c *Config) Validate() error {
	return nil
}
