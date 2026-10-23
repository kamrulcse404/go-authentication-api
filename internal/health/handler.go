package health

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "ok",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("failed to encode health response: %v", err)
	}
}
