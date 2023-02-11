package http

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/response"
)

var tracer = otel.Tracer("github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http")

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
	response.SetLogger(params.Logger)
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", params.Port),
		Handler: xray.Handler(xray.NewFixedSegmentNamer("app"), router),
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
