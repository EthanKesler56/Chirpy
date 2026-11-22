package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) getAllChirps(w http.ResponseWriter, r *http.Request) {
	allChirps := []chirp{}

	dbAllChirps, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		log.Printf("Error with chirp: %v", err)
		w.WriteHeader(500)
		return
	}
	for _, dbChirp := range dbAllChirps {
		allChirps = append(allChirps, chirp{ID: dbChirp.ID, Created_at: dbChirp.CreatedAt, Updated_at: dbChirp.UpdatedAt,
			Body: dbChirp.Body, User_id: dbChirp.UserID},
		)
	}
	data, err := json.Marshal(allChirps)
	if err != nil {
		log.Printf("Error retreiving chirps: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}
