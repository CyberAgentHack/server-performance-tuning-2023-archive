package database

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
)

type ViewingHistory struct {
}

func NewViewingHistory() *ViewingHistory {
	return &ViewingHistory{}
}

func (e *ViewingHistory) Create(ctx context.Context, viewingHistory *entity.ViewingHistory) (*entity.ViewingHistory, error) {
	_, span := tracer.Start(ctx, "database.ViewingHistory#Create")
	defer span.End()
	// TODO
	return &entity.ViewingHistory{ID: "id"}, nil
}

func (e *ViewingHistory) BatchGet(ctx context.Context, ids []string, userID string) (entity.ViewingHistories, error) {
	_, span := tracer.Start(ctx, "database.ViewingHistory#BatchGet")
	defer span.End()
	// TODO
	return entity.ViewingHistories{{ID: "id"}}, nil
}
