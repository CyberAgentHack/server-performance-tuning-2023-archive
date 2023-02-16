package http

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
)

const (
	keyCacheControl = "Cache-Control"
	keyContentType  = "Content-Type"
	keyVary         = "Vary"
)

const (
	applicationJSON = "application/json"
)

func (s *Service) send(status int, msg interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set(keyVary, "accept,accept-encoding,origin")
	if w.Header().Get(keyCacheControl) == "" {
		w.Header().Set(keyCacheControl, "private,no-store")
	}

	if msg == nil {
		w.WriteHeader(status)
		return
	}

	w.Header().Set(keyContentType, applicationJSON)
	w.WriteHeader(status)
	data, err := json.Marshal(msg)
	if err != nil {
		s.Error(err, w, r)
		return
	}
	w.Write(data)
}

func (s *Service) Error(err error, w http.ResponseWriter, r *http.Request) {
	const internal = 500
	status := httpStatus(err)
	if status >= internal {
		s.logger.Error("Error occurred", zap.Error(err))
	}
	s.sendError(err, status, w, r)
}

func (s *Service) BadRequest(err error, w http.ResponseWriter, r *http.Request) {
	s.sendError(err, http.StatusBadRequest, w, r)
}

func (s *Service) sendError(err error, status int, w http.ResponseWriter, r *http.Request) {
	s.send(status, err.Error(), w, r)
}

func httpStatus(err error) int {
	return errcode.HTTPStatus(err)
}

func (s *Service) OK(msg interface{}, w http.ResponseWriter, r *http.Request) {
	s.send(http.StatusOK, msg, w, r)
}
