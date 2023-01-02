package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
)

func TestCreateViewingHistoryRequestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		req          *CreateViewingHistoryRequest
		expectedCode errcode.Code
	}{
		{
			name:         "viewingHistory is nil",
			req:          &CreateViewingHistoryRequest{},
			expectedCode: errcode.CodeInvalidArgument,
		},
		{
			name: "success",
			req: &CreateViewingHistoryRequest{
				ViewingHistory: &entity.ViewingHistory{ID: "id"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.req.validate()
			require.Equal(t, tt.expectedCode, errcode.GetCode(actual))
		})
	}
}

func TestUsecaseImpl_CreateViewingHistory(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tests := []struct {
		name         string
		setup        func(*mocks)
		req          *CreateViewingHistoryRequest
		expected     *CreateViewingHistoryResponse
		expectedCode errcode.Code
	}{
		{
			name:         "validate error",
			setup:        func(m *mocks) {},
			req:          &CreateViewingHistoryRequest{},
			expectedCode: errcode.CodeInvalidArgument,
		},
		{
			name: "failed to Create",
			setup: func(m *mocks) {
				m.viewingHistory.EXPECT().Create(gomock.Any(), &entity.ViewingHistory{ID: "id"}).
					Return(nil, errcode.NewInternal("error"))
			},
			req: &CreateViewingHistoryRequest{
				ViewingHistory: &entity.ViewingHistory{ID: "id"},
			},
			expectedCode: errcode.CodeInternal,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.viewingHistory.EXPECT().Create(gomock.Any(), &entity.ViewingHistory{ID: "id"}).
					Return(&entity.ViewingHistory{ID: "id"}, nil)
			},
			req: &CreateViewingHistoryRequest{
				ViewingHistory: &entity.ViewingHistory{ID: "id"},
			},
			expected: &CreateViewingHistoryResponse{
				ViewingHistory: &entity.ViewingHistory{ID: "id"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			u := newUsecase(m)
			actual, err := u.CreateViewingHistory(ctx, tt.req)
			require.Equal(t, tt.expectedCode, errcode.GetCode(err))
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestUsecaseImpl_BatchGetViewingHistories(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tests := []struct {
		name         string
		setup        func(*mocks)
		req          *BatchGetViewingHistoriesRequest
		expected     *BatchGetViewingHistoriesResponse
		expectedCode errcode.Code
	}{
		{
			name:         "validate error",
			setup:        func(m *mocks) {},
			req:          &BatchGetViewingHistoriesRequest{},
			expectedCode: errcode.CodeInvalidArgument,
		},
		{
			name: "failed to BatchGet",
			setup: func(m *mocks) {
				m.viewingHistory.EXPECT().BatchGet(gomock.Any(), []string{"id"}, "userID").
					Return(nil, errcode.NewInternal("error"))
			},
			req: &BatchGetViewingHistoriesRequest{
				UserID:     "userID",
				EpisodeIDs: []string{"id"},
			},
			expectedCode: errcode.CodeInternal,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.viewingHistory.EXPECT().BatchGet(gomock.Any(), []string{"id"}, "userID").
					Return(entity.ViewingHistories{{ID: "id"}}, nil)
			},
			req: &BatchGetViewingHistoriesRequest{
				UserID:     "userID",
				EpisodeIDs: []string{"id"},
			},
			expected: &BatchGetViewingHistoriesResponse{
				ViewingHistories: entity.ViewingHistories{{ID: "id"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			u := newUsecase(m)
			actual, err := u.BatchGetViewingHistories(ctx, tt.req)
			require.Equal(t, tt.expectedCode, errcode.GetCode(err))
			require.Equal(t, tt.expected, actual)
		})
	}
}
