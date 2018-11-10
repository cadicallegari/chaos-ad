package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"cadicallegari/chaos-ad/pkg/server"
	"cadicallegari/chaos-ad/pkg/storage"
)

var (
	// version is set at build time
	Version = "No version provided at build time"
)

func main() {

	store, err := storage.New()
	if err != nil {
		log.Fatal(err)
	}

	handler := server.New(store)

	port := "8080"
	server := &http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     handler,
		ReadTimeout: time.Minute,
	}

	fmt.Printf("Starting server listening port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}
