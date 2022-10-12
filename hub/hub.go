package hub

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/shoot_and_win/command"
	"github.com/shoot_and_win/player"
)

func NewHub(commands chan<- command.Command) *Hub {
	return &Hub{
		clients:  make(map[string]*websocket.Conn),
		commands: commands,
	}
}

type Hub struct {
	clients  map[string]*websocket.Conn
	commands chan<- command.Command
}

func (h *Hub) AddClient(name string, conn *websocket.Conn) {
	h.clients[name] = conn
}

func (h *Hub) Read(p player.Player) {
	go func() {
		for {
			_, msg, err := p.Conn.ReadMessage()
			if err != nil {
				log.Printf("failed to read message from %s: %v\n", p.Name, err)
				continue
			}

			var cmd Command
			err = json.Unmarshal(msg, &cmd)
			if err != nil {
				log.Printf("failed to unmarshal message from %s: %v\n", p.Name, err)
				continue
			}

			cmd.payload = msg

			h.commands <- cmd
		}
	}()
}

func (h *Hub) Write(p player.Player, event []byte) error {
	return p.Conn.WriteMessage(websocket.TextMessage, event)
}

type Command struct {
	Command string `json:"command"`
	payload []byte
}

func (c Command) Name() string {
	return c.Command
}

func (c Command) Payload() []byte {
	return c.payload
}
