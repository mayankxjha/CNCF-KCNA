package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go Server!")
}

func main() {
	http.HandleFunc("/", helloHandler) // Register the handler for the root path

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil)) // Start the server
}
