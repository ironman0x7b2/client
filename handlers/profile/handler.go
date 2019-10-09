package profile

import (
	"io/ioutil"
	"net/http"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

func ProfilePicHandlerFunc(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseMultipartForm(10 << 20)
		file, _, err := r.FormFile("myFile")
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "Error Retrieving the File",
				Info:    err.Error(),
			})
			return
		}

		defer file.Close()

		tempFile, err := ioutil.TempFile("profile-pics", "upload-*.png")
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "Error Retrieving the File",
				Info:    err.Error(),
			})
			return
		}

		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "Error Retrieving the File",
				Info:    err.Error(),
			})
			return
		}

		tempFile.Write(fileBytes)
		utils.WriteResultToResponse(w, 200, "Successfully Uploaded File\n")
		return

	}
}
