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
		
		allPlayers, err := s.RegisterPlayer(body)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(allPlayers)
	}
}

func WebsocketHandler(h *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			panic(err)
		}

		h.AddClient(name, conn)
	}
}