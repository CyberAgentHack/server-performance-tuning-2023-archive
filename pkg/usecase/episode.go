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
	SeriesID string
}

type ListEpisodesResponse struct {
	Episodes entity.Episodes
	Series   entity.SeriesMulti
	Seasons  entity.Seasons
}

func (u *UsecaseImpl) ListEpisodes(ctx context.Context, req *ListEpisodesRequest) (*ListEpisodesResponse, error) {
	ctx, span := tracer.Start(ctx, "usecase.UsecaseImpl#ListEpisodes")
	defer span.End()

	params := &repository.ListEpisodesParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		SeasonID: req.SeasonID,
	}
	episodes, err := u.db.Episode.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	series := make(entity.SeriesMulti, 0, len(episodes))
	for i := range episodes {
		l, err := u.db.Series.List(ctx, &repository.ListSeriesParams{Limit: 1, SeriesID: episodes[i].SeriesID})
		if err != nil {
			return nil, errcode.New(err)
		}
		if len(l) == 0 {
			continue
		}
		series = append(series, l[0])
	}

	seasons := make(entity.Seasons, 0, len(episodes))
	for i := range episodes {
		if episodes[i].SeasonID == "" {
			continue
		}
		l, err := u.db.Season.List(ctx, &repository.ListSeasonsParams{Limit: 1, SeasonID: episodes[i].SeasonID})
		if err != nil {
			return nil, errcode.New(err)
		}
		if len(l) == 0 {
			continue
		}
		seasons = append(seasons, l[0])
	}

	return &ListEpisodesResponse{
		Episodes: episodes,
		Series:   series,
		Seasons:  seasons,
	}, nil
}
