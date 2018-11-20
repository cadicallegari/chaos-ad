package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cadicallegari/chaos-ad/pkg/cache"
	"cadicallegari/chaos-ad/pkg/server"
)

var (
	// version is set at build time
	Version = "No version provided at build time"
)

func main() {

	ttl, err := time.ParseDuration(os.Getenv("CACHE_TTL"))
	if err != nil {
		log.Fatal(err)
	}

	// cache, err := cache.NewLocal()
	cache, err := cache.NewRedis(
		os.Getenv("REDIS_URL"),
		"",
		0,
	)
	if err != nil {
		log.Fatal(err)
	}

	handler := server.New(cache, ttl)

	port := "8080"
	server := &http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     handler,
		ReadTimeout: time.Minute,
	}

	fmt.Printf("Starting server listening port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}
