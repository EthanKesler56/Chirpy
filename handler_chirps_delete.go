package main

import (
	"chirpy/internal/auth"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpDelete(w http.ResponseWriter, r *http.Request) {
	type returnvals struct {
		Error string `json:"error"`
	}

	//Getting Token
	token, err := auth.GetBearerToken(r.Header)
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

	//Validating Token
	validate_token, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		log.Printf("Error validating token %v", err)
		vals := returnvals{Error: "Something went wrong"}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(401)
		data, err_three := json.Marshal(vals)
		if err_three != nil {
			log.Print(err_three)

		}
		w.Write(data)
		return
	}
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
	if validate_token != uuid_chirp.UserID {
		http.Error(w, "unauthorized action", http.StatusForbidden)
		return
	}
	err = cfg.db.DeleteChirp(r.Context(), uuid_chirp.ID)
	if err != nil {
		log.Printf("Error deleting chirp %v", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}
