package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type ListSeriesRequest struct {
	Limit  int
	Offset int
}

type ListSeriesResponse struct {
	Series entity.SeriesMulti
	Casts  entity.Casts
}

func (u *UsecaseImpl) ListSeries(ctx context.Context, req *ListSeriesRequest) (*ListSeriesResponse, error) {
	const key = "series"
	params := &repository.ListSeriesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	do := func() (interface{}, error) {
		return u.db.Series.List(ctx, params)
	}
	v, err, _ := u.group.Do(key, do)
	if err != nil {
		return nil, errcode.New(err)
	}

	series := v.(entity.SeriesMulti)
	casts := make(entity.Casts, 0, len(series))
	for i := range series {
		c, err := u.db.Cast.BatchGet(ctx, series[i].CastIDs)
		if err != nil {
			return nil, errcode.New(err)
		}
		casts = append(casts, c...)
	}
	return &ListSeriesResponse{Series: series, Casts: casts}, nil
}
