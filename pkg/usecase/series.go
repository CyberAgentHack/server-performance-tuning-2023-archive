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
	Genres entity.Genres
}

func (u *UsecaseImpl) ListSeries(ctx context.Context, req *ListSeriesRequest) (*ListSeriesResponse, error) {
	ctx, span := tracer.Start(ctx, "usecase.UsecaseImpl#ListSeries")
	defer span.End()

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
	genreIDs := series.GenreIDs()
	genres, err := u.db.Genre.BatchGet(ctx, genreIDs)
	if err != nil {
		return nil, errcode.New(err)
	}
	return &ListSeriesResponse{Series: series, Genres: genres}, nil
}
