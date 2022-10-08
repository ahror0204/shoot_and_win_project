package main

import (
	"fmt"
	"strconv"
)

type Player struct {
	Name   string `json:"name"`
	Health uint16 `json:"health"`
}

type ReceiveMessage struct {
	Command    string `json:"command"`
	PlayerName string `json:"player_name"`
}

type SendMessage struct {
	Event      string `json:"event"`
	PlayerName string `json:"player_name"`
}

func (p *Player) GetShot() string {
	p.Health -= 10

	if p.Health == 0 {
		return fmt.Sprintf("Player:%s \nHealth status: died", p.Name)
	}

	return strconv.Itoa(int(p.Health))
}
