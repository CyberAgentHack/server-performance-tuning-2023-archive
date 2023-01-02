package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type ListSeriesRequest struct {
	PageSize int
}

type ListSeriesResponse struct {
	Series entity.SeriesMulti
}

func (u *UsecaseImpl) ListSeries(ctx context.Context, req *ListSeriesRequest) (*ListSeriesResponse, error) {
	params := &repository.ListSeriesParams{
		PageSize: req.PageSize,
	}
	series, err := u.db.Series.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	// TODO エピソード集計

	return &ListSeriesResponse{Series: series}, nil
}
