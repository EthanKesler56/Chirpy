package main

import (
	"chirpy/internal/database"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

////Creating CHIRP Section

type chirp struct {
	ID         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Body       string    `json:"body"`
	User_id    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) createChirp(w http.ResponseWriter, r *http.Request) {

	///DECODING SECTION

	type returnvals struct {
		Error string `json:"error"`
	}

	decoder := json.NewDecoder(r.Body)
	params := chirp{}
	err := decoder.Decode(&params)
	vals := returnvals{Error: "Something went wrong"}
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(400)
		data, err := json.Marshal(vals)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(data)
		return
	}

	if len(params.Body) > 140 {
		vals = returnvals{Error: "Chirp is too long"}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(400)
		data, err := json.Marshal(vals)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(data)
		return
	}
	chirpParams := database.CreateChirpParams{
		Body:   params.Body,
		UserID: params.User_id,
	}

	dbChirp, err := cfg.db.CreateChirp(r.Context(), chirpParams)
	if err != nil {
		log.Printf("Error with chirp: %v", err)
		w.WriteHeader(500)
		return
	}
	resp := chirp{
		ID:         dbChirp.ID,
		Created_at: dbChirp.CreatedAt,
		Updated_at: dbChirp.UpdatedAt,
		Body:       dbChirp.Body,
		User_id:    dbChirp.UserID,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal failed %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(data)

}
