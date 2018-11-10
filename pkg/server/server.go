package server

import (
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"net/http"
)

type serv struct {
	router *http.ServeMux
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

func (s *serv) handlePostProductsRequest(w http.ResponseWriter, req *http.Request) {

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

func New() *http.ServeMux {
	s := serv{
		router: http.NewServeMux(),
	}

	s.routes()

	return s.router
}
