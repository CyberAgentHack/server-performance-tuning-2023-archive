package database

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/ent"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Episode struct {
}

func NewEpisode() *Episode {
	return &Episode{}
}

func (e *Episode) GetCount(ctx context.Context, id string) (int, error) {
	// TODO
	return 10, nil
}

func (e *Episode) List(ctx context.Context, params *repository.ListEpisodesParams) (ent.Episodes, error) {
	// TODO
	return ent.Episodes{{ID: 1}}, nil
}
