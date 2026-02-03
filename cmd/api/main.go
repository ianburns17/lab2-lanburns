package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ianburns17/lab2-Ianburns/internal/routes"
)

func main() {
	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("starting server on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
