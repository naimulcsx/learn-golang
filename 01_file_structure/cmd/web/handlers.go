package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to Go Web Server</h1>")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Set status code
	w.WriteHeader(http.StatusOK)

	// Write response
	fmt.Fprintf(w, `{"status":"running","version":"1.0.0"}`)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Get the id parameter from the URL
	id := r.PathValue("id")

	// Write response
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "User ID: %s", id)
}