package main

import (
	"chirpy/internal/auth"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUpgradeUser(w http.ResponseWriter, r *http.Request) {

	getApiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		http.Error(w, "Couldn't get apiKey", http.StatusUnauthorized)
		return
	}
	if getApiKey != cfg.polkaSecret {
		http.Error(w, "incorrect API key", http.StatusUnauthorized)
		return
	}

	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		}
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(204)
		return
	}

	_, err = cfg.db.UpgradeUser(r.Context(), params.Data.UserID)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}
