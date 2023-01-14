package database

import (
	"context"
	"database/sql"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Season struct {
	db *sql.DB
}

func NewSeason(db *sql.DB) *Season {
	return &Season{db: db}
}

func (e *Season) List(ctx context.Context, params *repository.ListSeasonsParams) (entity.Seasons, error) {
	// TODO
	return entity.Seasons{{ID: "seasonID"}}, nil
}
