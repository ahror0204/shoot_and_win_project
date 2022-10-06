package main
import (
	"github.com/gorilla/websocket"
)
type Hub struct {
	clients map[string]*websocket.Conn
	makeShoot <-chan []MakeShoot
}
