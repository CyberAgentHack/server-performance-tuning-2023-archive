package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/ent"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/util/errcode"
)

type ListSeasonsRequest struct {
	PageSize int
	SeriesID string
}

type ListSeasonsResponse struct {
	Seasons ent.Seasons
}

func (u *UsecaseImpl) ListSeasons(ctx context.Context, req *ListSeasonsRequest) (*ListSeasonsResponse, error) {
	params := &repository.ListSeasonsParams{
		PageSize: req.PageSize,
		SeriesID: req.SeriesID,
	}
	episodes, err := u.db.Season.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	return &ListSeasonsResponse{Seasons: episodes}, nil
}
