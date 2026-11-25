package main

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handler_login(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
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
	getUser, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		log.Printf("failed to find user")
		http.Error(w, "Incorrect email or password", http.StatusUnauthorized)
		return
	}
	match, errTwo := auth.CheckPasswordHash(params.Password, getUser.HashedPassword)
	if errTwo != nil || !match {
		http.Error(w, "Incorrect email or password", http.StatusUnauthorized)
		return

	}
	token, err := auth.MakeJWT(getUser.ID, cfg.secret, time.Hour)
	if err != nil {
		http.Error(w, "couldn't create error", http.StatusInternalServerError)
		return
	}

	refreshToken := auth.MakeRefreshToken()
	_, err = cfg.db.RefreshToken(r.Context(), database.RefreshTokenParams{Token: refreshToken, UserID: getUser.ID, ExpiresAt: time.Now().Add(time.Hour * 1440)})
	if err != nil {
		log.Printf("refresh token error: %v", err)
		w.WriteHeader(500)
		return
	}

	user := User{ID: uuid.UUID(getUser.ID), CreatedAt: getUser.CreatedAt, UpdatedAt: getUser.UpdatedAt, Email: getUser.Email, Token: token, RefreshToken: refreshToken}
	data, err := json.Marshal(user)
	if err != nil {
		log.Printf("Marshal failed %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}
