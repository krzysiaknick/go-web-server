package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello-world", helloWorldHandler)
	http.HandleFunc("/hello-world-html", helloWorldHtmlHandler)
	http.HandleFunc("/hello-world-json", helloWorldJSONHandler)
	http.HandleFunc("/syllabus", syllabusHandler)

	// Define the port to listen on
	port := ":8080"
	fmt.Println("Server is listening on port", port)

	// Start the server and log any errors
	go func() {
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()

	fmt.Println("For help accessing websites type 'help': ")
	for {

		var input string
		fmt.Scanln(&input)
		if input == "help" {
			displayHelp()

		} else {
			fmt.Println("Invalid input. Please type 'help'.")
		}
	}

}

func displayHelp() {
	fmt.Println("Go Web Server Help:")
	fmt.Println("Available Endpoints:")
	fmt.Println("/hello-world - http://localhost:8080/hello-world")
	fmt.Println("/hello-world-html - http://localhost:8080/hello-world-html")
	fmt.Println("/hello-world-json - http://localhost:8080/hello-world-json")
	fmt.Println("/syllabus - http://localhost:8080/syllabus")
}
