package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/request"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/response"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeSeason(r chi.Router) {
	r.Get("/", s.listSeasons)
}

func (s *Service) listSeasons(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, span := tracer.Start(ctx, "http.Service#listSeasons")
	defer span.End()

	req := &usecase.ListSeasonsRequest{
		Limit:    request.QueryIntDefault(r, "limit", 20),
		Offset:   request.QueryInt(r, "offset"),
		SeriesID: request.Query(r, "seriesId"),
	}
	resp, err := s.usecase.ListSeasons(ctx, req)
	if err != nil {
		response.Error(err, w, r)
		return
	}

	response.OK(&entity.ListSeasonsResponse{Seasons: resp.Seasons, Genres: resp.Genres}, w, r)
}
