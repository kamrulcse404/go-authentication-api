package httpserver

import (
	"authentication-api/internal/health"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(healthHandler *health.Handler) http.Handler {
	router := chi.NewRouter()
	router.Get("/health", healthHandler.Check)

	return router
}
