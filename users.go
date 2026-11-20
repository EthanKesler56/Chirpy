package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"uuid"
)

// CREATING USER SECTION
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) users(w http.ResponseWriter, r *http.Request) {

	///DECODING SECTION

	type parameters struct {
		Email string `json:"email"`
	}
	type returnErr struct {
		Error string `json:"error"`
	}
	vals := returnErr{Error: "Something went wrong"}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		data, err := json.Marshal(vals)
		if err != nil {
			log.Printf("marshal failed %v", err)
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(data)
		return
	}
	// END OF DECODING SECTION

	//CREATING USER SECTION
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
	//END OF CREATING USER SECTION

}
