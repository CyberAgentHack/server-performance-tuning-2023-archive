package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeSeries(r chi.Router) {
	r.Get("/", s.listSeries)
}

func (s *Service) listSeries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, span := startTrace(ctx, r, "http.Service#listSeries")
	defer span.End()

	req := &usecase.ListSeriesRequest{
		Limit:  QueryIntDefault(r, "limit", 20),
		Offset: QueryInt(r, "offset"),
	}
	resp, err := s.usecase.ListSeries(ctx, req)
	if err != nil {
		s.Error(err, w, r)
		return
	}

	w.Header().Set("cache-control", "public,max-age=0,s-maxage=1")
	s.OK(&entity.ListSeriesMultiResponse{SeriesMulti: resp.Series, Genres: resp.Genres}, w, r)
}
