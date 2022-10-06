package main

import "strconv"

type MakeShoot struct {
	PahPah string `json:"pah-pah"`
}

type Player struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Health uint16 `json:"health"`
}

func (p *Player) GetShot() string {
	p.Health -= 10

	if p.Health == 0 {
		return "The player died"
	}

	return strconv.Itoa(int(p.Health))
}