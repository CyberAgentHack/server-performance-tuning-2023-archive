package database

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Series struct {
}

func NewSeries() *Series {
	return &Series{}
}

func (e *Series) List(ctx context.Context, params *repository.ListSeriesParams) (entity.SeriesMulti, error) {
	// TODO
	return entity.SeriesMulti{{ID: "seriesID"}}, nil
}
