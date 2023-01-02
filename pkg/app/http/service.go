package http

import (
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

type Service struct {
	now     func() time.Time
	logger  *zap.Logger
	usecase usecase.Usecase
}

func NewService(usecase usecase.Usecase, logger *zap.Logger) *Service {
	v := &Service{
		now:     time.Now,
		logger:  logger,
		usecase: usecase,
	}
	return v
}

func (s *Service) Register(mux *chi.Mux) {
	mux.Mount("/", s.newRouter())
}

func (s *Service) newRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", livenessCheck)
	r.Route("/series", s.routeSeries)
	r.Route("/seasons", s.routeSeason)
	r.Route("/episodes", s.routeEpisode)
	r.Route("/viewingHistories", s.routeViewingHistory)
	return r
}