package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
)

const (
	applicationJSON = "application/json"
)

var logger *zap.Logger

func SetLogger(l *zap.Logger) {
	logger = l
}

type Option func(*options)

type options struct {
	cachable bool
	smaxage  time.Duration
}

// setHeaders sets cache control and surrogate keys.
func setHeaders(h http.Header, opts ...Option) {
	dopt := &options{}
	for i := range opts {
		opts[i](dopt)
	}
	// cache control
	value := "private,no-store"
	if dopt.cachable {
		const format = "public,max-age=0,s-maxage=%.0f"
		value = fmt.Sprintf(format, dopt.smaxage.Seconds())
	}
	h.Set("vary", "accept,accept-encoding,origin")
	h.Set("cache-control", value)
}

func Send(status int, msg interface{}, w http.ResponseWriter, r *http.Request, opts ...Option) {
	// write common headers
	setHeaders(w.Header(), opts...)

	// nil handling
	if msg == nil {
		w.WriteHeader(status)
		return
	}

	w.Header().Set("content-type", applicationJSON)
	w.WriteHeader(status)
	data, err := json.Marshal(msg)
	if err != nil {
		Error(err, w, r)
		return
	}
	if _, err = w.Write(data); err != nil {
		logger.Error("Failed to send", zap.Any("body", msg), zap.Error(err))
	}
}

func Error(err error, w http.ResponseWriter, r *http.Request) {
	const internal = 500
	status := httpStatus(err)
	if status >= internal {
		logger.Error("Error occurred", zap.Error(err))
	}
	sendError(err, status, w, r)
}

func sendError(err error, status int, w http.ResponseWriter, r *http.Request) {
	Send(status, err.Error(), w, r)
}

func BadRequest(err error, w http.ResponseWriter, r *http.Request) {
	sendError(err, http.StatusBadRequest, w, r)
}

func NotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func httpStatus(err error) int {
	return errcode.HTTPStatus(err)
}

func OK(msg interface{}, w http.ResponseWriter, r *http.Request, opts ...Option) {
	Send(http.StatusOK, msg, w, r, opts...)
}
