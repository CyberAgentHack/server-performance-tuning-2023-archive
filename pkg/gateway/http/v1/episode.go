package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/request"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/response"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeEpisode(r chi.Router) {
	r.Get("/", s.listEpisodes)
}

func (s *Service) listEpisodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &usecase.ListEpisodesRequest{
		PageSize: request.QueryIntDefault(r, "pageSize", 20),
		SeriesID: request.Query(r, "seriesId"),
		SeasonID: request.Query(r, "seasonId"),
	}
	resp, err := s.usecase.ListEpisodes(ctx, req)
	if err != nil {
		response.Error(err, w, r)
		return
	}

	response.OK(resp.Episodes, w, r)
}
