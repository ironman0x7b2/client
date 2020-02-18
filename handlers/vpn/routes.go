package vpn

import (
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, _cli *cli.CLI) {
	r.Name("StartSubscription").
		Methods("POST").Path("/subscription").
		HandlerFunc(startSubscription(_cli))
}
