package handler

import (
	"encoding/json"
	"net/http"
	"github.com/shoot_and_win/service"
	"github.com/shoot_and_win/player"
	"github.com/shoot_and_win/hub"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func RegisterPlayerHandler(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body player.Player
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

func WebsocketHandler(h *hub.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			panic(err)
		}

		h.AddClient(name, conn)
	}
}
