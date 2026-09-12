package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Server  string `json:"server"`
	Message string `json:"message"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	serverID := os.Getenv("SERVER_ID")
	if serverID == "" {
		serverID = "backend-local"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Server:  serverID,
			Message: "request processed",
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response := HealthResponse{
			Status: "ok",
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	address := ":" + port

	log.Printf(
		"starting demo backend id=%s address=%s",
		serverID,
		address,
	)

	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatal(err)
	}
}