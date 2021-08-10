package http

import (
	"encoding/json"
	"net/http"
)

func RespondClientError(rw http.ResponseWriter, err error) {
	RespondClientMultiError(rw, []error{err})
}

func RespondClientMultiError(rw http.ResponseWriter, errs []error) {
	var strErrs []string
	for _, e := range errs {
		strErrs = append(strErrs, e.Error())
	}

	resp := map[string]interface{}{
		"errors": strErrs,
	}
	json, _ := json.Marshal(resp)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write(json)
}

func RespondSuccess(rw http.ResponseWriter, body interface{}) {
	resp := map[string]interface{}{
		"data": body,
	}
	json, _ := json.Marshal(resp)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(json)
}
