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
	Genre          Genre
	ViewingHistory ViewingHistory
}

type Episode interface {
	List(ctx context.Context, params *ListEpisodesParams) (entity.Episodes, error)
}

type ListEpisodesParams struct {
	Limit    int
	Offset   int
	SeasonID string
	SeriesID string
}

type Series interface {
	Get(ctx context.Context, id string) (*entity.Series, error)
	List(ctx context.Context, params *ListSeriesParams) (entity.SeriesMulti, error)
	BatchGet(ctx context.Context, ids []string) (entity.SeriesMulti, error)
}

type ListSeriesParams struct {
	Limit    int
	Offset   int
	SeriesID string
}

type Season interface {
	Get(ctx context.Context, id string) (*entity.Season, error)
	List(ctx context.Context, params *ListSeasonsParams) (entity.Seasons, error)
	BatchGet(ctx context.Context, ids []string) (entity.Seasons, error)
}

type ListSeasonsParams struct {
	Limit    int
	Offset   int
	SeasonID string
	SeriesID string
}

type ViewingHistory interface {
	Create(ctx context.Context, viewingHistory *entity.ViewingHistory) (*entity.ViewingHistory, error)
	BatchGet(ctx context.Context, ids []string, userID string) (entity.ViewingHistories, error)
}

type Genre interface {
	BatchGet(ctx context.Context, ids []string) (entity.Genres, error)
}
