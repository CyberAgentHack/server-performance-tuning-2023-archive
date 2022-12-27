package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/response"
)

func (s *Service) routeSeries(r chi.Router) {
	r.Get("/", s.listSeries)
}

func (s *Service) listSeries(w http.ResponseWriter, r *http.Request) {
	response.OK(nil, w, r)
}
