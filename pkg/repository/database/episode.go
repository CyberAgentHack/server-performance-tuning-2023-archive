package database

import (
	"context"
	"database/sql"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Episode struct {
	db *sql.DB
}

func NewEpisode(db *sql.DB) *Episode {
	return &Episode{
		db: db,
	}
}

func (e *Episode) List(ctx context.Context, params *repository.ListEpisodesParams) (entity.Episodes, error) {
	ctx, span := tracer.Start(ctx, "database.Episode#List")
	defer span.End()

	args := make([]any, 0, 3)
	query := "SELECT * FROM episodes"
	if params.SeasonID != "" {
		query += " WHERE seasonId = ?"
		args = append(args, params.SeasonID)
	}
	query += " LIMIT ? OFFSET ?"
	args = append(args, params.Limit, params.Offset)

	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errcode.New(err)
	}

	var episodes entity.Episodes
	for rows.Next() {
		episode := &entity.Episode{}
		err = rows.Scan(
			&episode.ID,
			&episode.DisplayName,
			&episode.Description,
			&episode.ImageURL,
			&episode.GenreIDs,
			&episode.SeasonID,
			&episode.PublishStartTime,
			&episode.DisplayOrder,
		)
		if err != nil {
			break
		}

		episodes = append(episodes, episode)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errcode.New(closeErr)
	}
	return episodes, nil
}
