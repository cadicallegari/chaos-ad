// +build integration

package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// "github.com/go-redis/cache"
	// "github.com/go-redis/redis"
	// "github.com/vmihailenco/msgpack"

	"cadicallegari/chaos-ad/pkg/server"
	"cadicallegari/chaos-ad/pkg/storage"
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

// func setup(t *testing.T) (*http.ServeMux, func()) {
// )

// func main() {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     os.Getenv("REDIS_URL"),
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	pong, err := client.Ping().Result()
// 	fmt.Println(pong, err)
// }

// db := newDB(t)
// return server.New(db), func() {
// }
// }

func TestShouldBeHealth(t *testing.T) {
	store, _ := storage.New()
	srv := server.New(store)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/healthz", nil)
	ok(t, err, "creating healthz request")

	srv.ServeHTTP(res, req)
	equals(t, res.Code, http.StatusOK, "response code")
}

func TestHandleNewRecordProperly(t *testing.T) {
	store, _ := storage.New()
	srv := server.New(store)

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
