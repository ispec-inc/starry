package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response jsonを返す
func Response(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("error: %v\n", err)
	}
}

// Text テキストを返す
func Text(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}
