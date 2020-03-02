package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ironman0x7b2/client/types"
)

func WriteErrorToResponse(w http.ResponseWriter, code int, err *types.Error) {
	res := types.Response{
		Success: false,
		Error:   err,
	}

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Panicln(err)
	}
}

func WriteResultToResponse(w http.ResponseWriter, code int, result interface{}) {
	res := types.Response{
		Success: true,
		Result:  result,
	}

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}
