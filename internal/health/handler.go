package health

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusService interface {
	GetStatus() string
}

type Handler struct {
	service StatusService
}

func NewHandler(service StatusService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": h.service.GetStatus(),
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("failed to encode health response: %v", err)
	}
}
