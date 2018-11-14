package server

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"time"

	"cadicallegari/chaos-ad/pkg/storage"
)

const (
	// pass it ass paramter or env var
	cacheTTL = time.Minute * 10
)

type serv struct {
	storage *storage.Storage
	router  *http.ServeMux
}

func (s *serv) handleHealthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok")
	}

}

func (s *serv) handleProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			s.handlePostProductsRequest(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok")
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

	ok, err := s.storage.CheckCache(hash, cacheTTL)
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
	fmt.Sprintf("Error: %s", msg)
	fmt.Println(msg)
	http.Error(w, msg, statusCode)
}

func New(store *storage.Storage) *http.ServeMux {
	s := serv{
		storage: store,
		router:  http.NewServeMux(),
	}

	s.routes()

	return s.router
}
