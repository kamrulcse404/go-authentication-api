package main

import (
	"authentication-api/internal/config"
	"authentication-api/internal/health"
	httpserver "authentication-api/internal/http"
	"authentication-api/internal/platform/database"
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()
	dbPool, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer dbPool.Close()

	log.Println("Database connection established")

	healthService := health.NewService()
	healthHandler := health.NewHandler(healthService)

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
