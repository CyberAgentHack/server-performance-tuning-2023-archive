package database

import (
	"context"
	"database/sql"
	"strings"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
)

type Cast struct {
	db *sql.DB
}

func NewCast(db *sql.DB) *Cast {
	return &Cast{db: db}
}

func (c *Cast) Get(ctx context.Context, id string) (*entity.Cast, error) {
	var cast *entity.Cast
	err := c.db.QueryRowContext(ctx, "SELECT * FROM casts WHERE id = ?", id).
		Scan(cast)
	if err != nil {
		return nil, errcode.New(err)
	}
	return cast, nil
}

func (c *Cast) BatchGet(ctx context.Context, ids []string) (entity.Casts, error) {
	rows, err := c.db.QueryContext(ctx, `SELECT * FROM casts WHERE id IN (?`+strings.Repeat(",?", len(ids)-1)+`)`, ids)
	if err != nil {
		return nil, errcode.New(err)
	}

	var casts entity.Casts
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			break
		}

		casts = append(casts, &entity.Cast{ID: id})
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errcode.New(closeErr)
	}

	return casts, nil
}
