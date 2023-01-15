package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type ListEpisodesRequest struct {
	Limit    int
	Offset   int
	SeasonID string
}

type ListEpisodesResponse struct {
	Episodes entity.Episodes
	Casts    entity.Casts
}

func (u *UsecaseImpl) ListEpisodes(ctx context.Context, req *ListEpisodesRequest) (*ListEpisodesResponse, error) {
	params := &repository.ListEpisodesParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		SeasonID: req.SeasonID,
	}
	episodes, err := u.db.Episode.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	castIDs := episodes.CastIDs()
	casts, err := u.db.Cast.BatchGet(ctx, castIDs)
	if err != nil {
		return nil, errcode.New(err)
	}
	return &ListEpisodesResponse{Episodes: episodes, Casts: casts}, nil
}
