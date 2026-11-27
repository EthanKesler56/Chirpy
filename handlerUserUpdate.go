package main

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"encoding/json"
	"log"
	"time"

	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUserUpdate(w http.ResponseWriter, r *http.Request) {

	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		Id          uuid.UUID `json:"id"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Email       string    `json:"email"`
		IsChirpyRed bool      `json:"is_chirpy_red"`
	}

	getToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "error getting token")
		return
	}
	UserID, err := auth.ValidateJWT(getToken, cfg.secret)
	if err != nil {

		respondWithError(w, http.StatusUnauthorized, "invalid token")
		return
	}
	decoder := json.NewDecoder(r.Body)
	params := credentials{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error")
		return
	}

	hashPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	//RAW RESPONSE
	updateUser, err := cfg.db.UpdateUser(r.Context(), database.UpdateUserParams{Email: params.Email, HashedPassword: hashPassword, ID: UserID})
	if err != nil {
		log.Printf("error updating user: %v", err)
		w.WriteHeader(500)
		return
	}
	//RAW RESPONSE

	resp := response{Id: updateUser.ID, CreatedAt: updateUser.CreatedAt, UpdatedAt: updateUser.UpdatedAt, Email: updateUser.Email, IsChirpyRed: updateUser.IsChirpyRed}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal failed:%v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}
