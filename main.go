package main

import (
	"log"
	"net/http"

	"github.com/AbdelazizTina-dev/go-credit-card-validator/handlers"
)

func main() {
	http.HandleFunc("/validate", handlers.ValidateHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
