package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func TestListEpisodes(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(m *mocks)
		expected     *entity.ListEpisodesResponse
		expectedCode int
	}{
		{
			name: "failed to ListEpisodes",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListEpisodes(gomock.Any(), &usecase.ListEpisodesRequest{
					Limit:    10,
					SeasonID: "seasonId",
				}).Return(nil, errcode.NewInternal("error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListEpisodes(gomock.Any(), &usecase.ListEpisodesRequest{
					Limit:    10,
					SeasonID: "seasonId",
				}).Return(&usecase.ListEpisodesResponse{
					Episodes: entity.Episodes{{ID: "episodeId"}},
					Series:   entity.SeriesMulti{{ID: "seriesId"}},
					Seasons:  entity.Seasons{{ID: "seasonsID"}},
				}, nil)
			},
			expected: &entity.ListEpisodesResponse{
				Episodes: entity.Episodes{{ID: "episodeId"}},
				Series:   entity.SeriesMulti{{ID: "seriesId"}},
				Seasons:  entity.Seasons{{ID: "seasonsID"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/episodes?limit=10&seasonId=seasonId", nil)
			newMux(m).ServeHTTP(w, r)
			res := w.Result()
			if res.StatusCode != http.StatusOK {
				require.Equal(t, tt.expectedCode, res.StatusCode)
				return
			}
			ret := &entity.ListEpisodesResponse{}
			require.NoError(t, json.NewDecoder(w.Body).Decode(&ret))
			require.Equal(t, tt.expected, ret)
		})
	}
}
