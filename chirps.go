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
	//resp := database.CreateChirpParams{
	//	Body:   params.Body,
	//	UserID: params.User_id,
	//}
	resp := chirp{
		ID: params.ID,
		Created_at: params.Created_at,
		Updated_at: params.Updated_at,
		Body: params.Body,
		User_id: params.User_id,
	}




	dbChirp, err := cfg.db.CreateChirp(r.Context(), resp)
	if err != nil {
		log.Printf("Error with chirp: %v", err)
		w.WriteHeader(500)
		return
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

/*

	user, err := cfg.db.CreateUser(r.Context(), params.Email)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		w.WriteHeader(500)
		return
	}
	resp := User{
		ID:        uuid.UUID(user.ID),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
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
*/
