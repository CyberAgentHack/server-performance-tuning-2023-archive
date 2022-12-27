package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type API interface {
	Register(mux *chi.Mux)
}

type RunServerParams struct {
	Group  *errgroup.Group
	Port   int
	Logger *zap.Logger
	APIs   []API
}

func RunServer(params *RunServerParams) (*http.Server, error) {
	const compressLevel = 6
	router := chi.NewRouter()
	router.Use(
		newCORS(),
		middleware.Compress(compressLevel),
	)
	return runServer(router, params.Group, params.Port, params.APIs...), nil
}

func runServer(router *chi.Mux, group *errgroup.Group, port int, apis ...API) *http.Server {
	setDefaultHandlers(router)
	for i := range apis {
		apis[i].Register(router)
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	serve := func() error {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	group.Go(serve)
	return server
}

func setDefaultHandlers(mux *chi.Mux) {
	mux.NotFound(notFoundHandler)
	mux.Get(HealthCheckPath, healthCheck)
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
			"X-API-Key",
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
