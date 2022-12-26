package http

import (
	"net/http"
)

func livenessCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
