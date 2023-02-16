package http

import (
	"net/http"
	"strconv"
	"strings"
)

func Query(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func QueryStrings(r *http.Request, key string) []string {
	q := r.URL.Query().Get(key)
	if q == "" {
		return nil
	}
	return strings.Split(q, ",")
}

func QueryInt64(r *http.Request, key string) int64 {
	v := r.URL.Query().Get(key)
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func QueryInt64Default(r *http.Request, key string, def int64) int64 {
	v := QueryInt64(r, key)
	if v == 0 {
		return def
	}
	return v
}

func QueryInt(r *http.Request, key string) int {
	return int(QueryInt64(r, key))
}

func QueryIntDefault(r *http.Request, key string, def int) int {
	return int(QueryInt64Default(r, key, int64(def)))
}
