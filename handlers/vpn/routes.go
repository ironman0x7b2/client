package vpn

import (
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, _cli *cli.CLI) {
	r.Name("StartSubscription").
		Methods("POST").Path("/subscription").
		HandlerFunc(startSubscriptionHandler(_cli))
	r.Name("SubmitSubscription").
		Methods("POST").Path("/subscription/submit").
		HandlerFunc(submitTxHashToNodeHandler())
	r.Name("GetVPNKey").
		Methods("GET").Path("/subscription/{id}/key").
		HandlerFunc(getVPNKeyHandler())
	r.Name("ConnectVPN").
		Methods("POST").Path("/subscription/{id}/websocket").
		HandlerFunc(connectVPNHandler())
	r.Name("GetSubscriptions").
		Methods("Get").Path("/subscriptions/{address}").
		HandlerFunc(getSubscriptionsHandler(_cli))
}
