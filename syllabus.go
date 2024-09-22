package main

import (
	_ "embed"
	"fmt"

	"net/http"
)

//go:embed gws.json
var syllabusJSON []byte

func syllabusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.Write(syllabusJSON)
	case http.MethodDelete:
		fmt.Fprintf(w, "deleted - stubbed")
	case http.MethodPost:
		fmt.Fprintf(w, "create - stubbed")
	case http.MethodPut:
		fmt.Fprintf(w, "update - stubbed")
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
