package usecase

import (
	"context"
	"fmt"

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
	key := fmt.Sprintf("series:%d:%d", req.Limit, req.Offset)
	params := &repository.ListSeriesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	do := func() (any, error) {
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
