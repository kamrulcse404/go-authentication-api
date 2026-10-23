package main

import (
	"authentication-api/internal/config"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.Load()
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"status": "ok",
		}

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("failed to encode health response: %v", err)
		}
	})

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
