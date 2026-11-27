package main

import (
	"chirpy/internal/database"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getAllChirps(w http.ResponseWriter, r *http.Request) {
	authorIDStr := r.URL.Query().Get("author_id")
	var author_uuid uuid.UUID
	var err error
	var dbChirps []database.Chirp
	allChirps := []chirp{}
	if authorIDStr != "" {
		author_uuid, err = uuid.Parse(authorIDStr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}
	}
	if author_uuid != uuid.Nil {
		dbChirps, err = cfg.db.GetChirpsByAuthor(r.Context(), author_uuid)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Bad Request")
		}

	} else {
		dbChirps, err = cfg.db.GetAllChirps(r.Context())
	}
	for _, dbChirp := range dbChirps {
		allChirps = append(allChirps, chirp{ID: dbChirp.ID, Created_at: dbChirp.CreatedAt, Updated_at: dbChirp.UpdatedAt,
			Body: dbChirp.Body, User_id: dbChirp.UserID},
		)
	}
	/*
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

	*/
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
