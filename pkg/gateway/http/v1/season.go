package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/response"
)

func (s *Service) routeSeason(r chi.Router) {
	r.Get("/", s.listSeason)
}

func (s *Service) listSeason(w http.ResponseWriter, r *http.Request) {
	response.OK(nil, w, r)
}
