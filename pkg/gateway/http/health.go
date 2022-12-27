package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

const HealthCheckPath = "/"

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type healthCheckServer struct{}

func NewHealthCheckServer() API {
	return &healthCheckServer{}
}

func (h *healthCheckServer) Register(mux *chi.Mux) {
	mux.Get(HealthCheckPath, healthCheck)
}
