package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/response"
)

func (s *Service) routeEpisode(r chi.Router) {
	r.Get("/", s.listEpisode)
}

func (s *Service) listEpisode(w http.ResponseWriter, r *http.Request) {
	response.OK(nil, w, r)
}
