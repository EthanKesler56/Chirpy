package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func validate_chirpy(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string `json:"body"`
	}
	type returnvals struct {
		Error string `json:"error"`
	}
	type successval struct {
		Success bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
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
	success := successval{Success: true}
	val, err := json.Marshal(success)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(val)

}

/*
  w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))`

*/
