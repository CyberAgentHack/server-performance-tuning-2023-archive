package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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

	fields := []string{
		"episodeID",
		"seasonID",
		"seriesID",
		"displayName",
		"description",
		"imageURL",
		"displayOrder",
	}

	clauses := make([]string, 0, 3)
	args := make([]any, 0, 3)
	if params.SeasonID != "" {
		clauses = append(clauses, "seasonID = ?")
		args = append(args, params.SeasonID)
	}
	if params.SeriesID != "" {
		clauses = append(clauses, "seriesID = ?")
		args = append(args, params.SeasonID)
	}

	var whereClause string
	if len(clauses) != 0 {
		whereClause = fmt.Sprintf("WHERE %s", strings.Join(clauses, " AND "))
	}

	query := fmt.Sprintf(
		"SELECT %s FROM episodes %s ORDER BY displayOrder LIMIT %d OFFSET %d",
		strings.Join(fields, ", "),
		whereClause,
		params.Limit,
		params.Offset,
	)

	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errcode.New(err)
	}

	var episodes entity.Episodes
	for rows.Next() {
		episode := &entity.Episode{}
		err = rows.Scan(
			&episode.ID,
			&episode.SeasonID,
			&episode.SeriesID,
			&episode.DisplayName,
			&episode.Description,
			&episode.ImageURL,
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
	if err != nil {
		return nil, errcode.New(err)
	}
	return episodes, nil
}
