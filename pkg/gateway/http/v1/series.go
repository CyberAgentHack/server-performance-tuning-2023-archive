package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/request"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/response"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeSeries(r chi.Router) {
	r.Get("/", s.listSeries)
}

func (s *Service) listSeries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &usecase.ListSeriesRequest{
		PageSize: request.QueryIntDefault(r, "pageSize", 20),
	}
	resp, err := s.usecase.ListSeries(ctx, req)
	if err != nil {
		response.Error(err, w, r)
		return
	}

	response.OK(resp.Series, w, r)
}
