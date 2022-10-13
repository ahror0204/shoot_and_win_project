package player

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	Name                 string `json:"name"`
	Health               uint16 `json:"health"`
	Conn                 *websocket.Conn
	isWaitingForOpponent bool
}

func (p *Player) SetWaitingForOpponent(b bool) {
	p.isWaitingForOpponent = b
}

func (p *Player) IsWaitingForOpponent() bool {
	return p.isWaitingForOpponent
}