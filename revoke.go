package main

import (
	"chirpy/internal/auth"
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) HandlerRevoke(w http.ResponseWriter, r *http.Request) {
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
	_, err = cfg.db.RevokeRefreshToken(r.Context(), bearer_token)
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
	w.WriteHeader(204)

}
