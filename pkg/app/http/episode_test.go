package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
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
		expected     entity.Episodes
		expectedCode int
	}{
		{
			name: "failed to ListEpisodes",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListEpisodes(gomock.Any(), &usecase.ListEpisodesRequest{
					PageSize: 10,
					SeriesID: "seriesId",
					SeasonID: "seasonId",
				}).Return(nil, errcode.NewInternal("error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListEpisodes(gomock.Any(), &usecase.ListEpisodesRequest{
					PageSize: 10,
					SeriesID: "seriesId",
					SeasonID: "seasonId",
				}).Return(&usecase.ListEpisodesResponse{
					Episodes: entity.Episodes{{ID: "id"}},
				}, nil)
			},
			expected: entity.Episodes{{ID: "id"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			u := newService(m)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			q := r.URL.Query()
			q.Add("pageSize", "10")
			q.Add("seriesId", "seriesId")
			q.Add("seasonId", "seasonId")
			r.URL.RawQuery = q.Encode()
			rctx := chi.NewRouteContext()
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			u.listEpisodes(w, r)
			res := w.Result()
			if res.StatusCode != http.StatusOK {
				require.Equal(t, tt.expectedCode, res.StatusCode)
				return
			}
			ret := entity.Episodes{}
			require.NoError(t, json.NewDecoder(w.Body).Decode(&ret))
			require.Equal(t, tt.expected, ret)
		})
	}
}