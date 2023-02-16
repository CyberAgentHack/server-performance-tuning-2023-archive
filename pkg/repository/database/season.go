package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
	ctx, span := tracer.Start(ctx, "database.Season#List")
	defer span.End()

	fields := []string{
		"seasonID",
		"seriesID",
		"displayName",
		"imageURL",
		"displayOrder",
	}

	clauses := make([]string, 0, 3)
	args := make([]any, 0, 3)
	if params.SeriesID != "" {
		clauses = append(clauses, "seriesID = ?")
		args = append(args, params.SeriesID)
	}
	if params.SeasonID != "" {
		clauses = append(clauses, "seasonID = ?")
		args = append(args, params.SeasonID)
	}

	var whereClause string
	if len(clauses) != 0 {
		whereClause = fmt.Sprintf("WHERE %s", strings.Join(clauses, " AND "))
	}

	query := fmt.Sprintf(
		"SELECT %s FROM seasons %s ORDER BY displayOrder LIMIT %d OFFSET %d",
		strings.Join(fields, ", "),
		whereClause,
		params.Limit,
		params.Offset,
	)

	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errcode.New(err)
	}

	var seasons entity.Seasons
	for rows.Next() {
		season := &entity.Season{}
		err = rows.Scan(
			&season.ID,
			&season.SeriesID,
			&season.DisplayName,
			&season.ImageURL,
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
	if err != nil {
		return nil, errcode.New(err)
	}
	return seasons, nil
}

func (e *Season) Get(ctx context.Context, id string) (*entity.Season, error) {
	ctx, span := tracer.Start(ctx, "database.Season#Get")
	defer span.End()

	fields := []string{
		"seasonID",
		"seriesID",
		"displayName",
		"imageURL",
		"displayOrder",
	}

	query := fmt.Sprintf(
		"SELECT %s FROM seasons WHERE seasonID = ?",
		strings.Join(fields, ", "),
	)
	row := e.db.QueryRowContext(ctx, query, id)
	if err := row.Err(); err != nil {
		return nil, errcode.New(err)
	}

	season := &entity.Season{}
	err := row.Scan(
		&season.ID,
		&season.SeriesID,
		&season.DisplayName,
		&season.ImageURL,
		&season.DisplayOrder,
	)
	return season, errcode.New(err)
}

func (e *Season) BatchGet(ctx context.Context, ids []string) (entity.Seasons, error) {
	ctx, span := tracer.Start(ctx, "database.Season#BatchGet")
	defer span.End()

	return nil, errcode.New(errors.New("not implemtented yet"))
}
