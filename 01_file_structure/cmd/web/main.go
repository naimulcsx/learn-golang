package main

import (
	"log"
	"net/http"
)

func main() {
    // Create a new ServeMux
    mux := http.NewServeMux()

    // Register route handlers
    mux.HandleFunc("/", homeHandler)
    mux.HandleFunc("/api/status", statusHandler)
    mux.HandleFunc("/users/{id}", userHandler)

    // Start the server
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal(err)
    }
}

