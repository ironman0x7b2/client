package tx

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/utils"
)

func getTx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		tx, err := cli.GetTx(vars["hash"])
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)
			return
		}

		utils.WriteResultToResponse(w, 200, tx)
	}
}
