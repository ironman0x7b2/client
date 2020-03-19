package vpn

import (
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/types"
)

func RegisterRoutes(r *mux.Router, _cli *cli.CLI, cfg *types.Config) {
	r.Name("GetResolversNodes").
		Methods("GET").Path("/nodes").
		HandlerFunc(getResolversNodesHandler(_cli, cfg))
	r.Name("GetResolverNodes").
		Methods("GET").Path("/nodes/{id}").
		HandlerFunc(getResolverNodesHandler(_cli, cfg))
	r.Name("GetSubscriptions").
		Methods("GET").Path("/subscriptions/{address}").
		HandlerFunc(getSubscriptionsHandler(_cli))
	r.Name("StartSubscription").
		Methods("POST").Path("/subscription").
		HandlerFunc(startSubscriptionHandler(_cli))
	r.Name("StartVPNConnection").
		Methods("POST").Path("/vpn").
		HandlerFunc(startVPNConnectionHandler(_cli))
	r.Name("EndVPNConnection").
		Methods("DELETE").Path("/vpn").
		HandlerFunc(endVPNConnectionHandler())
}
