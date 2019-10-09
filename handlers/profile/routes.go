package profile

import (
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
)

func RegisterRoutes(r *mux.Router, cli *_cli.CLI) {
	r.Name("Upload profile pic").
		Methods("POST").Path("/accounts/profilePic/upload").
		HandlerFunc(ProfilePicHandlerFunc(cli))
}
