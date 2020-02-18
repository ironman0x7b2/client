package types

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Resolver struct {
	ID   string `json:"id"`
	IP   string `json:"ip"`
	Port uint64 `json:"port"`
}

type Config struct {
	ChainID         string     `json:"chain_id"`
	RPCAddress      string     `json:"rpc_address"`
	ExplorerAddress string     `json:"explorer_address"`
	Resolvers       []Resolver `json:"resolvers"`
	VerifierDir     string     `json:"verifier_dir"`
	KeysDir         string     `json:"keys_dir"`

	uh func(nc *Config) error

	TrustNode  bool `json:"trust_node"`
	KillSwitch bool `json:"kill_switch"`
}

func NewDefaultConfig() *Config {
	return &Config{
		ChainID:         DefaultChainID,
		RPCAddress:      DefaultRPCAddress,
		ExplorerAddress: DefaultExplorerAddress,
		VerifierDir:     DefaultConfigDir,
		KeysDir:         DefaultConfigDir,
	}
}

func (c *Config) SetUpdateHook(h func(nc *Config) error) {
	c.uh = h
}

func (c *Config) UpdateHook(nc *Config) error {
	return c.uh(nc)
}

func (c *Config) Update(nc *Config) {
	if nc.ChainID != "" {
		c.ChainID = nc.ChainID
	}
	if nc.RPCAddress != "" {
		c.RPCAddress = nc.RPCAddress
	}
	if nc.ExplorerAddress != "" {
		c.ExplorerAddress = nc.ExplorerAddress
	}
	if nc.VerifierDir != "" {
		c.VerifierDir = nc.VerifierDir
	}
	if nc.KeysDir != "" {
		c.KeysDir = nc.KeysDir
	}
	if len(nc.Resolvers) > 0 {
		for _, resolver := range nc.Resolvers {
			c.Resolvers = append(c.Resolvers, resolver)
		}
	}

	c.TrustNode = nc.TrustNode
	c.KillSwitch = nc.KillSwitch
}

func (c *Config) LoadFromPath(path string) error {
	if path == "" {
		path = DefaultConfigFilePath
	}

	if _, err := os.Stat(path); err != nil {
		err = NewDefaultConfig().SaveToPath(path)
		if err != nil {
			return err
		}
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
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
	//if len(c.ResolverAddresess) == 0 {
	//	return errors.New("Required minimun one resolver node")
	//}

	return nil
}
