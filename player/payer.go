package player

import (
	"fmt"
	"strconv"

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

func (p *Player) GetShot() string {
	p.Health -= 10

	if p.Health == 0 {
		return fmt.Sprintf("Player:%s \nHealth status: died", p.Name)
	}

	return strconv.Itoa(int(p.Health))
}
