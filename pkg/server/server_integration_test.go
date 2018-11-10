// +build integration

package server_test

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

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

func setup(t *testing.T) (*http.ServeMux, func()) {
	db := newDB(t)
	return server.New(db), func() {
	}
}

func TestShouldBeHealth(t *testing.T) {
	srv := server.New(newDB(t))

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/healthz", nil)
	ok(t, err, "creating healthz request")

	srv.ServeHTTP(res, req)
	equals(t, res.Code, http.StatusOK, "response code")
}
