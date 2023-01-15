package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/request"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/response"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeSeries(r chi.Router) {
	r.Get("/", s.listSeries)
}

func (s *Service) listSeries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &usecase.ListSeriesRequest{
		Limit:  request.QueryIntDefault(r, "limit", 20),
		Offset: request.QueryInt(r, "offset"),
	}
	resp, err := s.usecase.ListSeries(ctx, req)
	if err != nil {
		response.Error(err, w, r)
		return
	}

	response.OK(&entity.ListSeriesMultiResponse{SeriesMulti: resp.Series, Casts: resp.Casts}, w, r)
}
