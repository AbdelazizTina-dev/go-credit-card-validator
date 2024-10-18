package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"unicode"

	"github.com/AbdelazizTina-dev/go-credit-card-validator/services"
)

type Request struct {
	Number string `json:"number"`
}

type Response struct {
	Valid bool `json:"valid"`
}

func isDigitsOnly(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	var req Request

	err = json.Unmarshal(body, &req)

	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	if len(req.Number) != 16 || !isDigitsOnly(req.Number) {
		http.Error(w, "Unvalid credit card number format, only accept 16 digits string.", http.StatusBadRequest)
		return
	}

	response := Response{Valid: services.Is_valid(req.Number)}

	json.NewEncoder(w).Encode(response)
}
