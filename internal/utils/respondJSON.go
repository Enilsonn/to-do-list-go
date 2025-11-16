package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("error to write JSON: %v", err)
	}
}

func RespondError(w http.ResponseWriter, status int, message any) {
	RespondJSON(w, status, map[string]any{
		"error": message,
	})
}
