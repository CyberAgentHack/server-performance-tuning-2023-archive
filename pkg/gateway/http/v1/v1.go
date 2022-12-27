package v1

import (
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Service struct {
	now    func() time.Time
	logger *zap.Logger
}

func NewService(logger *zap.Logger) *Service {
	v := &Service{
		now:    time.Now,
		logger: logger,
	}
	return v
}

func (s *Service) Register(mux *chi.Mux) {
	mux.Mount("/v1", s.newRouter())
}

func (s *Service) newRouter() chi.Router {
	r := chi.NewRouter()
	r.Route("/series", s.routeSeries)
	r.Route("/seasons", s.routeSeason)
	r.Route("/episodes", s.routeEpisode)
	return r
}
