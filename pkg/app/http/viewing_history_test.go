package http

import (
	"bytes"
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

func TestCreateViewingHistory(t *testing.T) {
	viewingHistoryID := "id"
	tests := []struct {
		name         string
		setup        func(m *mocks)
		body         string
		expected     *entity.ViewingHistory
		expectedCode int
	}{
		{
			name:         "failed to bind request",
			setup:        func(m *mocks) {},
			body:         `{"id":` + viewingHistoryID + `}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "failed to UpdateViewingHistory",
			setup: func(m *mocks) {
				m.uc.EXPECT().CreateViewingHistory(gomock.Any(), &usecase.CreateViewingHistoryRequest{
					ViewingHistory: &entity.ViewingHistory{ID: viewingHistoryID},
				}).Return(nil, errcode.NewInternal("error"))
			},
			body:         `{"id":"` + viewingHistoryID + `"}`,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.uc.EXPECT().CreateViewingHistory(gomock.Any(), &usecase.CreateViewingHistoryRequest{
					ViewingHistory: &entity.ViewingHistory{ID: viewingHistoryID},
				}).Return(&usecase.CreateViewingHistoryResponse{
					ViewingHistory: &entity.ViewingHistory{ID: viewingHistoryID},
				}, nil)
			},
			body:         `{"id":"` + viewingHistoryID + `"}`,
			expected:     &entity.ViewingHistory{ID: viewingHistoryID},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/viewingHistories", bytes.NewBuffer([]byte(tt.body)))
			r.Header.Set("Content-Type", "application/json")
			newMux(m).ServeHTTP(w, r)
			res := w.Result()
			if res.StatusCode != http.StatusOK {
				require.Equal(t, tt.expectedCode, res.StatusCode)
				return
			}
			ret := &entity.ViewingHistory{}
			require.NoError(t, json.NewDecoder(w.Body).Decode(ret))
			require.Equal(t, tt.expected, ret)
		})
	}
}

func TestListViewingHistories(t *testing.T) {
	viewingHistoryID, userID := "id", "userID"
	episodeIDs := []string{"episodeID"}
	tests := []struct {
		name         string
		setup        func(m *mocks)
		expected     *entity.ViewingHistories
		expectedCode int
	}{
		{
			name: "failed to BatchGetViewingHistories",
			setup: func(m *mocks) {
				m.uc.EXPECT().BatchGetViewingHistories(gomock.Any(), &usecase.BatchGetViewingHistoriesRequest{
					UserID:     userID,
					EpisodeIDs: episodeIDs,
				}).Return(nil, errcode.NewInternal("error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			setup: func(m *mocks) {
				m.uc.EXPECT().BatchGetViewingHistories(gomock.Any(), &usecase.BatchGetViewingHistoriesRequest{
					UserID:     userID,
					EpisodeIDs: episodeIDs,
				}).Return(&usecase.BatchGetViewingHistoriesResponse{
					ViewingHistories: entity.ViewingHistories{{ID: viewingHistoryID}},
				}, nil)
			},
			expected:     &entity.ViewingHistories{{ID: viewingHistoryID}},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			tt.setup(m)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/viewingHistories?episodeIds=episodeID", nil)
			r.Header.Set("userId", userID)
			newMux(m).ServeHTTP(w, r)
			res := w.Result()
			if res.StatusCode != http.StatusOK {
				require.Equal(t, tt.expectedCode, res.StatusCode)
				return
			}
			ret := &entity.ViewingHistories{}
			require.NoError(t, json.NewDecoder(w.Body).Decode(ret))
			require.Equal(t, tt.expected, ret)
		})
	}
}
