package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type ListEpisodesRequest struct {
	PageSize int
	SeriesID string
	SeasonID string
}

type ListEpisodesResponse struct {
	Episodes entity.Episodes
}

func (u *UsecaseImpl) ListEpisodes(ctx context.Context, req *ListEpisodesRequest) (*ListEpisodesResponse, error) {
	params := &repository.ListEpisodesParams{
		PageSize: req.PageSize,
		SeriesID: req.SeriesID,
		SeasonID: req.SeasonID,
	}
	episodes, err := u.db.Episode.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	return &ListEpisodesResponse{Episodes: episodes}, nil
}
