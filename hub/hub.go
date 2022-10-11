package hub

import (
	"github.com/gorilla/websocket"
)

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*websocket.Conn),
	}
}

type Hub struct {
	clients map[string]*websocket.Conn
}

func (h *Hub) Run() {

}

func (h *Hub) AddClient(name string, conn *websocket.Conn) {
	h.clients[name] = conn
}
