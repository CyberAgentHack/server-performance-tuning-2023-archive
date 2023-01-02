package usecase

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
	mock_repository "github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository/mock"
)

type mocks struct {
	episode        *mock_repository.MockEpisode
	series         *mock_repository.MockSeries
	season         *mock_repository.MockSeason
	viewingHistory *mock_repository.MockViewingHistory
}

func newMocks(t *testing.T) *mocks {
	ctrl := gomock.NewController(t)
	return &mocks{
		episode:        mock_repository.NewMockEpisode(ctrl),
		series:         mock_repository.NewMockSeries(ctrl),
		season:         mock_repository.NewMockSeason(ctrl),
		viewingHistory: mock_repository.NewMockViewingHistory(ctrl),
	}
}

func newUsecase(m *mocks) *UsecaseImpl {
	return &UsecaseImpl{
		validate: validator.New(),
		db: &repository.Database{
			Episode:        m.episode,
			Series:         m.series,
			Season:         m.season,
			ViewingHistory: m.viewingHistory,
		},
	}
}
