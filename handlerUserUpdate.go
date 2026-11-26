package main

import (
	"chirpy/internal/auth"

	"net/http"
)

func (cfg *apiConfig) handlerUserUpdate(w http.ResponseWriter, r *http.Request) {

	type returnvals struct {
		Error string `json:"error"`
	}

	getToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "error getting token")
	}
	user_id, err := auth.ValidateJWT(getToken, cfg.secret)
	if err != nil {

		respondWithError(w, http.StatusUnauthorized, "invalid token")
		return
	}
}
