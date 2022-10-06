package main

import (
	"encoding/json"
	"net/http"
)

func RegisterPlayerHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body Player
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err)
		}

		body.Health = 100
		
		err = s.RegisterPlayer(body)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}