package main

import (
	"authentication-api/internal/config"
	"authentication-api/internal/health"
	httpserver "authentication-api/internal/http"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.Load()
	healthHandler := health.NewHandler()

	router := httpserver.NewRouter(healthHandler)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Starting server on port %s", cfg.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
