package server

import (
	"log"
	"net/http"
)

func logRequestMidleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received:", r.Method, "at:", r.URL, "from:", r.RemoteAddr)
		fn(w, r)
	}
}

func (s *serv) routes() {
	s.router.HandleFunc("/healthz", logRequestMidleware(s.handleHealthz()))
	s.router.HandleFunc("/v1/products", logRequestMidleware(s.handleProducts()))
}
