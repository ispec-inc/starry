package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Response(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

func Text(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}
