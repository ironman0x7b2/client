package config

import (
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/client/types"
)

func RegisterRoutes(r *mux.Router, config *types.Config) {
	r.Name("GetConfig").
		Methods("GET").Path("/config").
		HandlerFunc(getConfigHandler(config))
	r.Name("UpdateConfig").
		Methods("PUT").Path("/config").
		HandlerFunc(updateConfigHandler(config))
}
