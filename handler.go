package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World – GWS")
}

func helloWorldHtmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hello World – GWS</h1>")
}

func helloWorldJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Hello World - GWS"}
	json.NewEncoder(w).Encode(response)
}

func d6(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(6) + 1

	response := map[string]int{"diceRoll": randomNum}
	json.NewEncoder(w).Encode(response)
}
