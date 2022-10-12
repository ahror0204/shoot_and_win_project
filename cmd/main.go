package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/shoot_and_win/command"
	"github.com/shoot_and_win/handler"
	"github.com/shoot_and_win/hub"
	"github.com/shoot_and_win/player"
	"github.com/shoot_and_win/repository"
	"github.com/shoot_and_win/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

var upgrader = websocket.Upgrader{}

func main() {
	r := chi.NewRouter()
	commands := make(chan command.Command)
	repo := repository.NewInMemoryRepository()
	s := service.NewService(repo)
	h := hub.NewHub(commands)
	websocketHandlar := handler.NewWebsocketHandler(s, h, commands)
	websocketHandlar.Run()

	r.Use(cors.AllowAll().Handler)
	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		var p player.Player
		p.Name = r.URL.Query().Get("name")

		serverWS(w, r, s, h, p)
	})

	http.ListenAndServe("localhost:8080", r)
}

func serverWS(w http.ResponseWriter, r *http.Request, s service.Service, h *hub.Hub, p player.Player) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	p.Conn = conn

	s.CreatePlayer(p)

	h.Read(p)

	availablePlayers, err := json.Marshal(s.AvailablePlayers())
	if err != nil {
		panic(err)
	}
	 err = h.Write(p, availablePlayers)
	 if err != nil {
		panic(err)
	}
}

type REquest struct {
	Name string `json:"name"`
}
