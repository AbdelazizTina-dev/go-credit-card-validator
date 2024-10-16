package main

import "net/http"

func main() {
	http.HandleFunc("/validate", handlers.ValidateHandler)
	http.ListenAndServe(":8080", nil)
}
