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
}

type ListEpisodesResponse struct {
	Episodes entity.Episodes
	Genres   entity.Genres
}

func (u *UsecaseImpl) ListEpisodes(ctx context.Context, req *ListEpisodesRequest) (*ListEpisodesResponse, error) {
	params := &repository.ListEpisodesParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		SeasonID: req.SeasonID,
	}
	episodes, err := u.db.Episode.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	genres := make(entity.Genres, 0, len(episodes))
	for i := range episodes {
		c, err := u.db.Genre.BatchGet(ctx, episodes[i].GenreIDs)
		if err != nil {
			return nil, errcode.New(err)
		}
		genres = append(genres, c...)
	}
	return &ListEpisodesResponse{Episodes: episodes, Genres: genres}, nil
}
