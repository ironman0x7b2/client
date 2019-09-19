package key

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/client/types"
)

var (
	_ types.Request = (*addKey)(nil)
	_ types.Request = (*deleteKey)(nil)
)

type addKey struct {
	Name          string `json:"name"`
	Mnemonic      string `json:"mnemonic"`
	Password      string `json:"password"`
	BIP39Password string `json:"bip_39_password"`
}

func newAddKey(r *http.Request) (*addKey, error) {
	var body addKey
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (a *addKey) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("invalid field name")
	}
	if a.Password == "" {
		return fmt.Errorf("invalid field password")
	}

	return nil
}

type deleteKey struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func newDeleteKey(r *http.Request) (*deleteKey, error) {
	var body deleteKey
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	body.Name = vars["name"]

	return &body, nil
}

func (a *deleteKey) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("invalid field name")
	}
	if a.Password == "" {
		return fmt.Errorf("invalid field password")
	}

	return nil
}
