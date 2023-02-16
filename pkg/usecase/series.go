package usecase

import (
	"context"
	"fmt"
	"time"

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
	Genres entity.Genres
}

func (u *UsecaseImpl) ListSeries(ctx context.Context, req *ListSeriesRequest) (*ListSeriesResponse, error) {
	ctx, span := tracer.Start(ctx, "usecase.UsecaseImpl#ListSeries")
	defer span.End()

	key := fmt.Sprintf("%v", req)
	resp := &ListSeriesResponse{}
	hit, err := u.redis.Get(ctx, key, resp)
	if err != nil {
		return nil, errcode.New(err)
	}
	if hit {
		return resp, nil
	}

	params := &repository.ListSeriesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	series, err := u.db.Series.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	genreIDs := series.GenreIDs()
	genres, err := u.db.Genre.BatchGet(ctx, genreIDs)
	if err != nil {
		return nil, errcode.New(err)
	}

	resp = &ListSeriesResponse{
		Series: series,
		Genres: genres,
	}
	u.redis.Set(ctx, key, resp, time.Second*10)
	return resp, nil
}
