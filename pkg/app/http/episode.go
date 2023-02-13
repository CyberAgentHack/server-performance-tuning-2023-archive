package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/request"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/response"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeEpisode(r chi.Router) {
	r.Get("/", s.listEpisodes)
}

func (s *Service) listEpisodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, span := startTrace(ctx, r, "http.Service#listEpisodes")
	defer span.End()

	req := &usecase.ListEpisodesRequest{
		Limit:    request.QueryIntDefault(r, "limit", 20),
		Offset:   request.QueryInt(r, "offset"),
		SeasonID: request.Query(r, "seasonId"),
		SeriesID: request.Query(r, "seriesId"),
	}
	resp, err := s.usecase.ListEpisodes(ctx, req)
	if err != nil {
		response.Error(err, w, r)
		return
	}

	response.OK(&entity.ListEpisodesResponse{Episodes: resp.Episodes, Series: resp.Series, Seasons: resp.Seasons}, w, r)
}
