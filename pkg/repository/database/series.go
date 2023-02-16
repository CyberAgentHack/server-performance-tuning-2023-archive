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

type Series struct {
	db *sql.DB
}

func NewSeries(db *sql.DB) *Series {
	return &Series{db: db}
}

func (e *Series) List(ctx context.Context, params *repository.ListSeriesParams) (entity.SeriesMulti, error) {
	ctx, span := tracer.Start(ctx, "database.Series#List")
	defer span.End()

	fields := []string{"seriesID", "displayName", "description", "imageURL", "genreID"}

	clauses := make([]string, 0, 3)
	args := make([]any, 0, 3)
	if params.SeriesID != "" {
		clauses = append(clauses, "seriesID = ?")
		args = append(args, params.SeriesID)
	}

	var whereClause string
	if len(clauses) != 0 {
		whereClause = fmt.Sprintf("WHERE %s", strings.Join(clauses, " AND "))
	}

	query := fmt.Sprintf(
		"SELECT %s FROM series %s LIMIT %d OFFSET %d",
		strings.Join(fields, ", "),
		whereClause,
		params.Limit,
		params.Offset,
	)
	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errcode.New(err)
	}

	var seriesMulti entity.SeriesMulti
	for rows.Next() {
		series := &entity.Series{}
		err = rows.Scan(
			&series.ID,
			&series.DisplayName,
			&series.Description,
			&series.ImageURL,
			&series.GenreID,
		)
		if err != nil {
			break
		}
		seriesMulti = append(seriesMulti, series)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errcode.New(closeErr)
	}
	if err != nil {
		return nil, errcode.New(err)
	}

	return seriesMulti, nil
}

func (e *Series) Get(ctx context.Context, id string) (*entity.Series, error) {
	ctx, span := tracer.Start(ctx, "database.Series#Get")
	defer span.End()

	query := "SELECT seriesID, displayName, description, imageURL, genreID FROM series WHERE seriesID = ?"
	row := e.db.QueryRowContext(ctx, query, id)
	if err := row.Err(); err != nil {
		return nil, errcode.New(err)
	}

	series := &entity.Series{}
	err := row.Scan(
		&series.ID,
		&series.DisplayName,
		&series.Description,
		&series.ImageURL,
		&series.GenreID,
	)
	return series, errcode.New(err)
}

func (e *Series) BatchGet(ctx context.Context, ids []string) (entity.SeriesMulti, error) {
	ctx, span := tracer.Start(ctx, "database.Series#BatchGet")
	defer span.End()

	return nil, errcode.New(errors.New("not implemtented yet"))
}
