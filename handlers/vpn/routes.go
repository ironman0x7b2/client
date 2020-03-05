package vpn

import (
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, _cli *cli.CLI) {
	r.Name("GetSubscriptions").
		Methods("GET").Path("/subscriptions/{address}").
		HandlerFunc(getSubscriptionsHandler(_cli))
	r.Name("VPNConnection").
		Methods("POST").Path("/connection").
		HandlerFunc(connectVPNHandler(_cli))
	r.Name("EndConnection").
		Methods("DELETE").Path("/connection").
		HandlerFunc(endVPNHandler())
}
