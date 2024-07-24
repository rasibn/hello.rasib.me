package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler handles HTTP requests to the "/hello" endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to text/plain
	w.Header().Set("Content-Type", "text/html")
	// Write "Hello, World!" to the response
	fmt.Fprintln(w, "<h1>Hello, World!</h1>")
}

func main() {
	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Register the helloHandler function to handle requests to the "/hello" path
	mux.HandleFunc("GET /", helloHandler)

	// Start the HTTP server on port 8080
	log.Fatal(http.ListenAndServe(":8089", mux))
}
