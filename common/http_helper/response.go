package http_helper

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func SendServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("Server Error."))
}

func SendInvalidResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte("Invalid Request."))
}

func SendInvalidResponseWithErrs(w http.ResponseWriter, errs url.Values) {
	err := map[string]interface{}{"validation_error": errs}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(err)
}

func SendResponse(w http.ResponseWriter, status int, data Response) {
	w.Header().Set("Content-Type", "application/json'")
	_ = json.NewEncoder(w).Encode(data)
}
