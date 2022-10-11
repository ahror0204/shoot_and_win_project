package service

import (
	"github.com/shoot_and_win/match"
	"github.com/shoot_and_win/player"
)

type Repository interface {
	GetMatch(id string) match.Match
	UpdateMatch(match match.Match)
	CreateMatch(match match.Match)
	GetPlayer(name string) player.Player
	RemovePlayer(name string)
	ListPlayers(name string) []player.Player
	RegisterPlayer(player player.Player) error
}
