package database

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Season struct {
}

func NewSeason() *Season {
	return &Season{}
}

func (e *Season) List(ctx context.Context, params *repository.ListSeasonsParams) (entity.Seasons, error) {
	// TODO
	return entity.Seasons{{ID: "seasonID"}}, nil
}
