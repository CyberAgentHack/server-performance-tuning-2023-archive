package database

import (
	"context"
	"database/sql"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Episode struct {
	DB *sql.DB
}

func NewEpisode(db *sql.DB) *Episode {
	return &Episode{
		DB: db,
	}
}

func (e *Episode) GetCount(ctx context.Context, id string) (int, error) {
	// TODO
	return 10, nil
}

func (e *Episode) List(ctx context.Context, params *repository.ListEpisodesParams) (entity.Episodes, error) {
	rows, err := e.DB.QueryContext(ctx, "select * from episodes")
	if err != nil {
		return nil, errcode.New(err)
	}

	var episodes entity.Episodes
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			break
		}

		episodes = append(episodes, &entity.Episode{ID: id})
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errcode.New(closeErr)
	}

	return episodes, nil
}
