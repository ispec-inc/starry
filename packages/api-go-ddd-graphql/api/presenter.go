package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("error: %v\n", err)
	}
}

func Text(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}
