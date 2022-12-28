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

	"github.com/CyberAgentHack/server-performance-tuning-2023/ent"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func TestListSeries(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(m *mocks)
		expected     ent.SeriesSlice
		expectedCode int
	}{
		{
			name: "failed to ListSeries",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListSeries(gomock.Any(), &usecase.ListSeriesRequest{
					PageSize: 10,
				}).Return(nil, errcode.NewInternal("error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.uc.EXPECT().ListSeries(gomock.Any(), &usecase.ListSeriesRequest{
					PageSize: 10,
				}).Return(&usecase.ListSeriesResponse{
					Series: ent.SeriesSlice{{ID: 1}},
				}, nil)
			},
			expected: ent.SeriesSlice{{ID: 1}},
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
			r.URL.RawQuery = q.Encode()
			rctx := chi.NewRouteContext()
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			u.listSeries(w, r)
			res := w.Result()
			if res.StatusCode != http.StatusOK {
				require.Equal(t, tt.expectedCode, res.StatusCode)
				return
			}
			ret := ent.SeriesSlice{}
			require.NoError(t, json.NewDecoder(w.Body).Decode(&ret))
			require.Equal(t, tt.expected, ret)
		})
	}
}
