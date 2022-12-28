package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/ent"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/util/errcode"
)

type ListSeriesRequest struct {
	PageSize int
}

type ListSeriesResponse struct {
	Series ent.SeriesSlice
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
