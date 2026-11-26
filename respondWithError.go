package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, err := json.Marshal(errorResponse{Error: msg})
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	w.Write(data)
}
