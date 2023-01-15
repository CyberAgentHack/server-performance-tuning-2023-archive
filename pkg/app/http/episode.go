package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/request"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/response"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeEpisode(r chi.Router) {
	r.Get("/", s.listEpisodes)
}

func (s *Service) listEpisodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &usecase.ListEpisodesRequest{
		Limit:    request.QueryIntDefault(r, "limit", 20),
		Offset:   request.QueryInt(r, "offset"),
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
