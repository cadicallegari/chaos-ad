package server

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"time"

	"cadicallegari/chaos-ad/pkg/cache"
)

type serv struct {
	cache    cache.CacherHitter
	router   *http.ServeMux
	cacheTTL time.Duration
}

func (s *serv) handleHealthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "ok")
		default:
			handleError(w, http.StatusBadRequest, nil)
		}
	}

}

func (s *serv) handleProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			s.handlePostProductsRequest(w, r)
		default:
			handleError(w, http.StatusMethodNotAllowed, nil)
		}
	}

}

func (s *serv) handlePostProductsRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, r.Body); err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	ok, err := s.cache.Hit(hash, s.cacheTTL)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleError(w http.ResponseWriter, statusCode int, err error) {
	var msg string
	if err != nil {
		msg = fmt.Sprintf(`{"error": %q}`, err)
	}
	fmt.Printf("Sending error status code: %d, msg: %s\n", statusCode, msg)
	http.Error(w, msg, statusCode)
}

func New(cache cache.CacherHitter, ttl time.Duration) *http.ServeMux {
	s := serv{
		cache:    cache,
		cacheTTL: ttl,
		router:   http.NewServeMux(),
	}

	s.routes()

	return s.router
}
