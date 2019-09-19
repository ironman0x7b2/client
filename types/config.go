package types

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sentinel-official/hub/app"
	hub "github.com/sentinel-official/hub/types"
)

type Config struct {
	ChainID              string `json:"chain_id"`
	RPCAddress           string `json:"rpc_address"`
	KeysDir              string `json:"keys_dir"`
	KeyName              string `json:"key_name"`
	Bech32PrefixAccAddr  string `json:"bech_32_prefix_acc_addr"`
	Bech32PrefixAccPub   string `json:"bech_32_prefix_acc_pub"`
	Bech32PrefixValAddr  string `json:"bech_32_prefix_val_addr"`
	Bech32PrefixValPub   string `json:"bech_32_prefix_val_pub"`
	Bech32PrefixConsAddr string `json:"bech_32_prefix_cons_addr"`
	Bech32PrefixConsPub  string `json:"bech_32_prefix_cons_pub"`

	uh func() error
}

func NewDefaultConfig() *Config {
	return &Config{
		ChainID:              "sentinel-turing-1",
		RPCAddress:           "127.0.0.1:26657",
		KeysDir:              app.DefaultCLIHome,
		KeyName:              "",
		Bech32PrefixAccAddr:  hub.Bech32PrefixAccAddr,
		Bech32PrefixAccPub:   hub.Bech32PrefixAccPub,
		Bech32PrefixValAddr:  hub.Bech32PrefixValAddr,
		Bech32PrefixValPub:   hub.Bech32PrefixValPub,
		Bech32PrefixConsAddr: hub.Bech32PrefixConsAddr,
		Bech32PrefixConsPub:  hub.Bech32PrefixConsPub,
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
	if cfg.Bech32PrefixAccAddr != "" {
		c.Bech32PrefixAccAddr = cfg.Bech32PrefixAccAddr
	}
	if cfg.Bech32PrefixAccPub != "" {
		c.Bech32PrefixAccPub = cfg.Bech32PrefixAccPub
	}
	if cfg.Bech32PrefixValAddr != "" {
		c.Bech32PrefixValAddr = cfg.Bech32PrefixValAddr
	}
	if cfg.Bech32PrefixValPub != "" {
		c.Bech32PrefixValPub = cfg.Bech32PrefixValPub
	}
	if cfg.Bech32PrefixConsAddr != "" {
		c.Bech32PrefixConsAddr = cfg.Bech32PrefixConsAddr
	}
	if cfg.Bech32PrefixConsPub != "" {
		c.Bech32PrefixConsPub = cfg.Bech32PrefixConsPub
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
