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
	if len(ids) == 0 {
		return nil, nil
	}
	ctx, span := tracer.Start(ctx, "database.Cast#BatchGet")
	defer span.End()
	rows, err := c.db.QueryContext(ctx, `SELECT genreID, displayName FROM genres WHERE genreID IN (?`+strings.Repeat(",?", len(ids)-1)+`)`, convertStringsToAnys(ids)...)
	if err != nil {
		return nil, errcode.New(err)
	}

	var genres entity.Genres
	for rows.Next() {
		var genre entity.Genre
		err = rows.Scan(&genre.ID, &genre.DisplayName)
		if err != nil {
			break
		}
		genres = append(genres, &genre)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errcode.New(closeErr)
	}
	if err != nil {
		return nil, errcode.New(err)
	}

	return genres, nil
}
