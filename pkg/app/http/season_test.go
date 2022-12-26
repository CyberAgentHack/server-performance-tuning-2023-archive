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

func TestListSeasons(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(m *mocks)
		expected     *entity.ListSeasonsResponse
		expectedCode int
	}{
		{
			name: "failed to ListSeasons",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListSeasons(gomock.Any(), &usecase.ListSeasonsRequest{
					Limit:    10,
					SeriesID: "seriesId",
				}).Return(nil, errcode.NewInternal("error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListSeasons(gomock.Any(), &usecase.ListSeasonsRequest{
					Limit:    10,
					SeriesID: "seriesId",
				}).Return(&usecase.ListSeasonsResponse{
					Seasons: entity.Seasons{{ID: "seasonId"}},
				}, nil)
			},
			expected: &entity.ListSeasonsResponse{
				Seasons: entity.Seasons{{ID: "seasonId"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/seasons?limit=10&seriesId=seriesId", nil)
			newMux(m).ServeHTTP(w, r)
			res := w.Result()
			if res.StatusCode != http.StatusOK {
				require.Equal(t, tt.expectedCode, res.StatusCode)
				return
			}
			ret := &entity.ListSeasonsResponse{}
			require.NoError(t, json.NewDecoder(w.Body).Decode(&ret))
			require.Equal(t, tt.expected, ret)
		})
	}
}
