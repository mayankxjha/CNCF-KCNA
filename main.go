package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go Server!\n")
	fmt.Fprintf(w, "env var is: %s", os.Getenv("VAR_1"))
	fmt.Fprintf(w, "secret var is: %s", os.Getenv("VAR_2"))
}

func main() {
	http.HandleFunc("/", helloHandler) // Register the handler for the root path

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil)) // Start the server
}
