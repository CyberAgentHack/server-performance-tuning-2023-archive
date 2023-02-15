//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./usecase.go -destination=./mock/usecase.go

package usecase

import (
	"context"

	"github.com/go-playground/validator"
	"go.opentelemetry.io/otel"
	"golang.org/x/sync/singleflight"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/db"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

var (
	_      Usecase = (*UsecaseImpl)(nil)
	tracer         = otel.Tracer("github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase")
)

type Usecase interface {
	ListEpisodes(ctx context.Context, req *ListEpisodesRequest) (*ListEpisodesResponse, error)

	ListSeries(ctx context.Context, req *ListSeriesRequest) (*ListSeriesResponse, error)

	ListSeasons(ctx context.Context, req *ListSeasonsRequest) (*ListSeasonsResponse, error)

	CreateViewingHistory(ctx context.Context, req *CreateViewingHistoryRequest) (*CreateViewingHistoryResponse, error)
	BatchGetViewingHistories(ctx context.Context, req *BatchGetViewingHistoriesRequest) (*BatchGetViewingHistoriesResponse, error)
}

type UsecaseImpl struct {
	db       *repository.Database
	redis    db.RedisClient
	validate *validator.Validate
	group    *singleflight.Group
}

func NewUsecase(db *repository.Database, redis db.RedisClient) *UsecaseImpl {
	return &UsecaseImpl{
		db:       db,
		redis:    redis,
		validate: validator.New(),
		group:    &singleflight.Group{},
	}
}
