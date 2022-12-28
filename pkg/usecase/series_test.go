package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/CyberAgentHack/server-performance-tuning-2023/ent"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/util/errcode"
)

func TestUsecaseImpl_ListSeries(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tests := []struct {
		name         string
		setup        func(*mocks)
		req          *ListSeriesRequest
		expected     *ListSeriesResponse
		expectedCode errcode.Code
	}{
		{
			name: "failed to List",
			setup: func(m *mocks) {
				m.series.EXPECT().List(gomock.Any(), &repository.ListSeriesParams{
					PageSize: 10,
				}).Return(nil, errcode.NewInternal("error"))
			},
			req: &ListSeriesRequest{
				PageSize: 10,
			},
			expectedCode: errcode.CodeInternal,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.series.EXPECT().List(gomock.Any(), &repository.ListSeriesParams{
					PageSize: 10,
				}).Return(ent.SeriesSlice{{ID: 1}}, nil)
			},
			req: &ListSeriesRequest{
				PageSize: 10,
			},
			expected: &ListSeriesResponse{
				Series: ent.SeriesSlice{{ID: 1}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			u := newUsecase(m)
			actual, err := u.ListSeries(ctx, tt.req)
			require.Equal(t, tt.expectedCode, errcode.GetCode(err))
			require.Equal(t, tt.expected, actual)
		})
	}
}
