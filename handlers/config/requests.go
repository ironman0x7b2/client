package config

import (
	"encoding/json"
	"net/http"

	"github.com/ironman0x7b2/client/types"
)

var (
	_ types.Request = (*updateConfig)(nil)
)

type updateConfig struct {
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
}

func newUpdateConfig(r *http.Request) (*updateConfig, error) {
	var body updateConfig
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (u *updateConfig) Validate() error {
	return nil
}
