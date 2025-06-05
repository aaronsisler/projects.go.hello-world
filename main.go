package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function for the "/hello" endpoint
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
