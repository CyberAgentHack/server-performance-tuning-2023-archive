package http

import (
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase/mock"
)

type mocks struct {
	uc *mock.MockUsecase
}

func newMocks(t *testing.T) *mocks {
	ctrl := gomock.NewController(t)
	return &mocks{uc: mock.NewMockUsecase(ctrl)}
}

func newService(m *mocks) *Service {
	s := &Service{
		usecase: m.uc,
		now:     time.Now,
		logger:  zap.NewNop(),
	}
	return s
}

func newMux(m *mocks) *chi.Mux {
	mux := chi.NewMux()
	newService(m).Register(mux)
	return mux
}
