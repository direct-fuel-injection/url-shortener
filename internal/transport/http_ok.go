package transport

import (
	"encoding/json"
	"net/http"
)

func RespondOK(data any, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func RedirectPermanent(url string, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}