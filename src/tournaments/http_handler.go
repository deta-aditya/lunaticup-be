package tournaments

import (
	"encoding/json"
	"net/http"

	httpUtils "github.com/deta-aditya/lunaticup-be/src/http"
)

func HandleCreate(rw http.ResponseWriter, r *http.Request) {
	var tr TournamentRaw
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		httpUtils.RespondClientError(rw, err)
		return
	}

	if errs := tr.IsValid(); len(errs) > 0 {
		httpUtils.RespondClientMultiError(rw, errs)
		return
	}

	httpUtils.RespondSuccess(rw, nil)
}
