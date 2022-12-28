package v1

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/response"
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
	response.SetLogger(s.logger)
	return s
}
