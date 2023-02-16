package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeSeason(r chi.Router) {
	r.Get("/", s.listSeasons)
}

func (s *Service) listSeasons(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, span := startTrace(ctx, r, "http.Service#listSeasons")
	defer span.End()

	req := &usecase.ListSeasonsRequest{
		Limit:    QueryIntDefault(r, "limit", 20),
		Offset:   QueryInt(r, "offset"),
		SeriesID: Query(r, "seriesId"),
	}
	resp, err := s.usecase.ListSeasons(ctx, req)
	if err != nil {
		s.Error(err, w, r)
		return
	}

	w.Header().Set("cache-control", "public,max-age=0,s-maxage=1")
	s.OK(&entity.ListSeasonsResponse{Seasons: resp.Seasons}, w, r)
}
