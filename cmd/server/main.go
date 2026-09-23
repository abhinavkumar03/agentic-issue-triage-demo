package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"status": "ok",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("server listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}