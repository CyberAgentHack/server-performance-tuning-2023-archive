package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type ListSeasonsRequest struct {
	Limit    int
	Offset   int
	SeriesID string
}

type ListSeasonsResponse struct {
	Seasons entity.Seasons
}

func (u *UsecaseImpl) ListSeasons(ctx context.Context, req *ListSeasonsRequest) (*ListSeasonsResponse, error) {
	params := &repository.ListSeasonsParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		SeriesID: req.SeriesID,
	}
	episodes, err := u.db.Season.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	return &ListSeasonsResponse{Seasons: episodes}, nil
}
