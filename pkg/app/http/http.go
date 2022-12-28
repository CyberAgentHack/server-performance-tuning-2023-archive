package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

type ServerParams struct {
	Port    int
	Logger  *zap.Logger
	Service *Service
}

func NewServer(params *ServerParams) *http.Server {
	const compressLevel = 6
	router := chi.NewRouter()
	router.Use(
		newCORS(),
		middleware.Compress(compressLevel),
	)
	setDefaultHandlers(router)
	params.Service.Register(router)
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", params.Port),
		Handler: router,
	}
}

func setDefaultHandlers(mux *chi.Mux) {
	mux.NotFound(notFoundHandler)
}

func newCORS() func(http.Handler) http.Handler {
	const maxAge = 604800
	opts := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Accept-Encoding",
			"Content-Type",
			"Authorization",
			"If-Modified-Since",
			"X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge:           maxAge,
	}
	return cors.Handler(opts)
}

func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
