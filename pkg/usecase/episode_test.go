package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/CyberAgentHack/server-performance-tuning-2023/ent"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

func TestUsecaseImpl_ListEpisodes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tests := []struct {
		name         string
		setup        func(*mocks)
		req          *ListEpisodesRequest
		expected     *ListEpisodesResponse
		expectedCode errcode.Code
	}{
		{
			name: "failed to List",
			setup: func(m *mocks) {
				m.episode.EXPECT().List(gomock.Any(), &repository.ListEpisodesParams{
					PageSize: 10,
				}).Return(nil, errcode.NewInternal("error"))
			},
			req: &ListEpisodesRequest{
				PageSize: 10,
			},
			expectedCode: errcode.CodeInternal,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.episode.EXPECT().List(gomock.Any(), &repository.ListEpisodesParams{
					PageSize: 10,
				}).Return(ent.Episodes{{ID: 1}}, nil)
			},
			req: &ListEpisodesRequest{
				PageSize: 10,
			},
			expected: &ListEpisodesResponse{
				Episodes: ent.Episodes{{ID: 1}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			u := newUsecase(m)
			actual, err := u.ListEpisodes(ctx, tt.req)
			require.Equal(t, tt.expectedCode, errcode.GetCode(err))
			require.Equal(t, tt.expected, actual)
		})
	}
}
