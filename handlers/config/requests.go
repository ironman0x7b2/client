package config

import (
	"encoding/json"
	"net/http"

	"github.com/ironman0x7b2/client/types"
)

type updateConfig struct {
	ChainID         string         `json:"chain_id"`
	RPCAddress      string         `json:"rpc_address"`
	ExplorerAddress string         `json:"explorer_address"`
	VerifierDir     string         `json:"verifier_dir"`
	KeysDir         string         `json:"keys_dir"`
	Resolver        types.Resolver `json:"resolver"`
	TrustNode       bool           `json:"trust_node"`
	KillSwitch      bool           `json:"kill_switch"`
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
