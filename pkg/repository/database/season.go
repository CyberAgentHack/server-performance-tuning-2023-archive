package database

import (
	"context"
	"database/sql"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Season struct {
	db *sql.DB
}

func NewSeason(db *sql.DB) *Season {
	return &Season{db: db}
}

func (e *Season) List(ctx context.Context, params *repository.ListSeasonsParams) (entity.Seasons, error) {
	args := make([]any, 0, 3)
	query := "SELECT * FROM seasons"
	if params.SeriesID != "" {
		query += " WHERE seriesId = ?"
		args = append(args, params.SeriesID)
	}
	query += " LIMIT ? OFFSET ?"
	args = append(args, params.Limit, params.Offset)

	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errcode.New(err)
	}

	var seasons entity.Seasons
	for rows.Next() {
		season := &entity.Season{}
		err = rows.Scan(
			&season.ID,
			&season.DisplayName,
			&season.ImageURL,
			&season.CastIDs,
			&season.SeriesID,
			&season.DisplayOrder,
		)
		if err != nil {
			break
		}

		seasons = append(seasons, season)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errcode.New(closeErr)
	}
	return seasons, nil
}
