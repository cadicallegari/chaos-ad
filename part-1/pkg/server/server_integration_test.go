// +build integration

package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"cadicallegari/chaos-ad/pkg/cache"
	"cadicallegari/chaos-ad/pkg/server"
)

func assert(tb testing.TB, condition bool, msg string) {
	if !condition {
		tb.Error(msg)
	}
}

func ok(tb testing.TB, err error, msg string) {
	if err != nil {
		tb.Errorf("Error not expected: %s\n", err)
	}
}

func equals(tb testing.TB, exp, act interface{}, msg string) {
	if exp != act {
		tb.Errorf("%s, not equals: expecting %v, but got: %v", msg, exp, act)
	}
}

func TestShouldBeHealth(t *testing.T) {
	store, _ := cache.NewLocal()
	srv := server.New(store, time.Minute)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/healthz", nil)
	ok(t, err, "creating healthz request")

	srv.ServeHTTP(res, req)
	equals(t, res.Code, http.StatusOK, "response code")
}

func TestHandleNewRecordProperly(t *testing.T) {
	store, _ := cache.NewLocal()
	srv := server.New(store, time.Minute)

	body := `[{"id": "123", "name": "mesa"}]`

	res := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/products",
		strings.NewReader(body),
	)

	srv.ServeHTTP(res, req)
	equals(t, http.StatusOK, res.Code, "first request status code")

	res = httptest.NewRecorder()
	req = httptest.NewRequest(
		http.MethodPost,
		"/v1/products",
		strings.NewReader(body),
	)
	srv.ServeHTTP(res, req)
	equals(t, http.StatusForbidden, res.Code, "second request status code")

}
