package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/request"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/response"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeSeason(r chi.Router) {
	r.Get("/", s.listSeasons)
}

func (s *Service) listSeasons(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &usecase.ListSeasonsRequest{
		PageSize: request.QueryIntDefault(r, "pageSize", 20),
		SeriesID: request.Query(r, "seriesId"),
	}
	resp, err := s.usecase.ListSeasons(ctx, req)
	if err != nil {
		response.Error(err, w, r)
		return
	}

	response.OK(resp.Seasons, w, r)
}
