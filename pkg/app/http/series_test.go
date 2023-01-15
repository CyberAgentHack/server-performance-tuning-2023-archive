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

func TestListSeries(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(m *mocks)
		expected     entity.SeriesMulti
		expectedCode int
	}{
		{
			name: "failed to ListSeries",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListSeries(gomock.Any(), &usecase.ListSeriesRequest{
					Limit: 10,
				}).Return(nil, errcode.NewInternal("error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListSeries(gomock.Any(), &usecase.ListSeriesRequest{
					Limit: 10,
				}).Return(&usecase.ListSeriesResponse{
					Series: entity.SeriesMulti{{ID: "id"}},
				}, nil)
			},
			expected: entity.SeriesMulti{{ID: "id"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/series?limit=10", nil)
			newMux(m).ServeHTTP(w, r)
			res := w.Result()
			if res.StatusCode != http.StatusOK {
				require.Equal(t, tt.expectedCode, res.StatusCode)
				return
			}
			ret := entity.SeriesMulti{}
			require.NoError(t, json.NewDecoder(w.Body).Decode(&ret))
			require.Equal(t, tt.expected, ret)
		})
	}
}
