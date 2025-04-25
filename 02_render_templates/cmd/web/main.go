package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register route handlers
	mux.HandleFunc("GET /", Home)
	mux.HandleFunc("GET /about", About)
	mux.HandleFunc("GET /status", Status)
	mux.HandleFunc("GET /users/{id}", User)

	// Start server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

