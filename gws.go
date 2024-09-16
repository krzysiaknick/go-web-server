package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler function that responds with "Hello, World!"
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Register the handler for the "/" route
	http.HandleFunc("/", helloWorldHandler)

	// Define the port to listen on
	port := ":8080"
	fmt.Println("Server is listening on port", port)

	// Start the server and log any errors
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
