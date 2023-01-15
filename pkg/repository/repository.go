//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./repository.go -destination=./mock/repository.go

package repository

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
)

type Database struct {
	Episode        Episode
	Series         Series
	Season         Season
	Cast           Cast
	ViewingHistory ViewingHistory
}

type Episode interface {
	GetCount(ctx context.Context, id string) (int, error)
	List(ctx context.Context, params *ListEpisodesParams) (entity.Episodes, error)
}

type ListEpisodesParams struct {
	Limit    int
	Offset   int
	SeriesID string
	SeasonID string
}

type Series interface {
	List(ctx context.Context, params *ListSeriesParams) (entity.SeriesMulti, error)
}

type ListSeriesParams struct {
	Limit  int
	Offset int
}

type Season interface {
	List(ctx context.Context, params *ListSeasonsParams) (entity.Seasons, error)
}

type ListSeasonsParams struct {
	Limit    int
	Offset   int
	SeriesID string
}

type ViewingHistory interface {
	Create(ctx context.Context, viewingHistory *entity.ViewingHistory) (*entity.ViewingHistory, error)
	BatchGet(ctx context.Context, ids []string, userID string) (entity.ViewingHistories, error)
}

type Cast interface {
	Get(ctx context.Context, id string) (*entity.Cast, error)
	Batch(ctx context.Context, ids []string) (entity.Casts, error)
}
