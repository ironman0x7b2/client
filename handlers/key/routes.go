package key

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("GetKeys").
		Methods("GET").Path("/keys").
		HandlerFunc(getKeysHandler(cli))
	r.Name("AddKey").
		Methods("POST").Path("/keys").
		HandlerFunc(addKeyHandler(cli))
	r.Name("DeleteKey").
		Methods("POST").Path("/keys/{name}").
		HandlerFunc(deleteKeyHandler(cli))
}