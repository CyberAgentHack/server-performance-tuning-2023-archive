//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./usecase.go -destination=./mock/usecase.go

package usecase

import (
	"context"

	"github.com/go-playground/validator"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

var _ Usecase = (*UsecaseImpl)(nil)

type Usecase interface {
	ListEpisodes(ctx context.Context, req *ListEpisodesRequest) (*ListEpisodesResponse, error)

	ListSeries(ctx context.Context, req *ListSeriesRequest) (*ListSeriesResponse, error)

	ListSeasons(ctx context.Context, req *ListSeasonsRequest) (*ListSeasonsResponse, error)

	CreateViewingHistory(ctx context.Context, req *CreateViewingHistoryRequest) (*CreateViewingHistoryResponse, error)
	BatchGetViewingHistories(ctx context.Context, req *BatchGetViewingHistoriesRequest) (*BatchGetViewingHistoriesResponse, error)
}

type UsecaseImpl struct {
	db       *repository.Database
	validate *validator.Validate
}

type Config struct {
	DB *repository.Database
}

func NewUsecase(cfg *Config) *UsecaseImpl {
	return &UsecaseImpl{
		db:       cfg.DB,
		validate: validator.New(),
	}
}
