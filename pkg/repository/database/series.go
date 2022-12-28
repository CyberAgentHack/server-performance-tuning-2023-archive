package database

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/ent"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Series struct {
}

func NewSeries() *Series {
	return &Series{}
}

func (e *Series) List(ctx context.Context, params *repository.ListSeriesParams) (ent.SeriesSlice, error) {
	// TODO
	return ent.SeriesSlice{{ID: 1}}, nil
}
