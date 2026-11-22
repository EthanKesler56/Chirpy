package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getSingleChirps(w http.ResponseWriter, r *http.Request) {

	url_path := r.PathValue("chirpID")
	parseUuid, err := uuid.Parse(url_path)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	uuid_chirp, err := cfg.db.SingleChirp(r.Context(), parseUuid)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusNotFound)
		return
	}
	chirp_ID := chirp{ID: uuid_chirp.ID, Created_at: uuid_chirp.CreatedAt,
		Updated_at: uuid_chirp.UpdatedAt, Body: uuid_chirp.Body, User_id: uuid_chirp.UserID}

	data, err := json.Marshal(chirp_ID)
	if err != nil {
		log.Printf("marshal failed %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}
