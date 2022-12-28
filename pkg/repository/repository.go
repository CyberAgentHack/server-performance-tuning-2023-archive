//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./repository.go -destination=./mock/repository.go

package repository

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/ent"
)

type Database struct {
	Episode Episode
	Series  Series
	Season  Season
}

type Episode interface {
	GetCount(ctx context.Context, id string) (int, error)
	List(ctx context.Context, params *ListEpisodesParams) ([]*ent.Episode, error)
}

type ListEpisodesParams struct {
	PageSize int
	SeriesID string
	SeasonID string
}

type Series interface {
	List(ctx context.Context, params *ListSeriesParams) ([]*ent.Series, error)
}

type ListSeriesParams struct {
	PageSize int
}

type Season interface {
	List(ctx context.Context, params *ListSeasonsParams) ([]*ent.Season, error)
}

type ListSeasonsParams struct {
	PageSize int
	SeriesID string
}
