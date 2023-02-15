package response

import (
	"encoding/json"
	"net/http"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
)

const (
	keyCacheControl = "cache-control"
	keyContentType  = "content-type"
	keyVary         = "vary"
)

const (
	applicationJSON = "application/json"
)

func Send(status int, msg interface{}, w http.ResponseWriter, r *http.Request) {
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
		Error(err, w, r)
		return
	}
	w.Write(data)
}

func Error(err error, w http.ResponseWriter, r *http.Request) {
	status := httpStatus(err)
	sendError(err, status, w, r)
}

func BadRequest(err error, w http.ResponseWriter, r *http.Request) {
	sendError(err, http.StatusBadRequest, w, r)
}

func sendError(err error, status int, w http.ResponseWriter, r *http.Request) {
	Send(status, err.Error(), w, r)
}

func httpStatus(err error) int {
	return errcode.HTTPStatus(err)
}

func OK(msg interface{}, w http.ResponseWriter, r *http.Request) {
	Send(http.StatusOK, msg, w, r)
}
