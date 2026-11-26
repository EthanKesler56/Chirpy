package main

import (
	"chirpy/internal/auth"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	type returnvals struct {
		Error string `json:"error"`
	}

	bearer_token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("Error getting token: %v", err)
		vals := returnvals{Error: "Something went wrong"}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(401)
		data, err_two := json.Marshal(vals)
		if err_two != nil {
			log.Print(err_two)
		}
		w.Write(data)
		return
	}
	//refresh_user, err := cfg.db.RefreshToken(r.Context(), bearer_token)
	refresh_user, err := cfg.db.GetUserFromRefreshToken(r.Context(), bearer_token)
	if err != nil {
		log.Printf("Error getting token: %v", err)
		vals := returnvals{Error: "Something went wrong"}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(401)
		data, err_two := json.Marshal(vals)
		if err_two != nil {
			log.Print(err_two)
		}
		w.Write(data)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	new_token, err := auth.MakeJWT(refresh_user.ID, cfg.secret, time.Hour)
	if err != nil {
		log.Printf("Error getting token: %v", err)
		vals := returnvals{Error: "Something went wrong"}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(401)
		data, err_two := json.Marshal(vals)
		if err_two != nil {
			log.Print(err_two)
		}
		w.Write(data)
		return
	}
	resp := response{Token: new_token}
	w.WriteHeader(200)
	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Marshal Failed %v", err)
		w.WriteHeader(500)
		return
	}
	w.Write(data)

}
