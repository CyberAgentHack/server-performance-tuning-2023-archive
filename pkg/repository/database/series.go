package database

import (
	"context"
	"database/sql"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Series struct {
	db *sql.DB
}

func NewSeries(db *sql.DB) *Series {
	return &Series{db: db}
}

func (e *Series) List(ctx context.Context, params *repository.ListSeriesParams) (entity.SeriesMulti, error) {
	// TODO
	return entity.SeriesMulti{{ID: "seriesID"}}, nil
}
