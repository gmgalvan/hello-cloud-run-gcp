package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// A tiny struct for the /healthz JSON response
type health struct {
	Status string `json:"status"`
}

func main() {
	mux := http.NewServeMux()

	// GET /              → “Hello, Cloud Run!”
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Cloud Run!")
	})

	// GET /healthz       → {"status":"ok"}
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(health{Status: "ok"})
	})

	// GET /time          → 2025-07-31T12:34:56Z (RFC 3339)
	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, time.Now().UTC().Format(time.RFC3339))
	})

	// GET /echo?msg=hi   → hi   (falls back to a friendly default)
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("msg")
		if msg == "" {
			msg = "Nothing to echo 🤷"
		}
		fmt.Fprint(w, msg)
	})

	// Honour $PORT if Cloud Run (or other hosting) sets it; default to 8080 locally
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
