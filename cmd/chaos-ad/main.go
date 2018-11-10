package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"cadicallegari/chaos-ad/pkg/server"
)

var (
	// version is set at build time
	Version = "No version provided at build time"
)

func main() {

	handler := server.New()

	port := "8080"
	server := &http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     handler,
		ReadTimeout: time.Minute,
	}

	fmt.Printf("Starting server listening port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}
