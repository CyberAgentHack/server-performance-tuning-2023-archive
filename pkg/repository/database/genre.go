package database

import (
	"context"
	"database/sql"
	"strings"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
)

type Genre struct {
	db *sql.DB
}

func NewGenre(db *sql.DB) *Genre {
	return &Genre{db: db}
}

func (c *Genre) BatchGet(ctx context.Context, ids []string) (entity.Genres, error) {
	ctx, span := tracer.Start(ctx, "database.Cast#BatchGet")
	defer span.End()
	rows, err := c.db.QueryContext(ctx, `SELECT * FROM genres WHERE id IN (?`+strings.Repeat(",?", len(ids)-1)+`)`, ids)
	if err != nil {
		return nil, errcode.New(err)
	}

	var genres entity.Genres
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			break
		}

		genres = append(genres, &entity.Genre{ID: id})
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errcode.New(closeErr)
	}

	return genres, nil
}
