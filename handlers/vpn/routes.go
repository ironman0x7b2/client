package vpn

import (
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, _cli *cli.CLI) {
	r.Name("VPNConnection").
		Methods("POST").Path("/connect/new").
		HandlerFunc(connectVPNHandler(_cli))
	r.Name("EndConnection").
		Methods("POST").Path("/connect/end").
		HandlerFunc(endVPNHandler())
}
