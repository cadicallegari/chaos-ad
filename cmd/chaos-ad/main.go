package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"cadicallegari/chaos-ad/pkg/cache"
	"cadicallegari/chaos-ad/pkg/server"
)

var (
	// version is set at build time
	Version = "No version provided at build time"
)

func main() {

	cache, err := cache.NewLocal()
	if err != nil {
		log.Fatal(err)
	}

	handler := server.New(cache)

	port := "8080"
	server := &http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     handler,
		ReadTimeout: time.Minute,
	}

	fmt.Printf("Starting server listening port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}
