package match

import "github.com/shoot_and_win/player"

type Match struct {
	ID           string
	Player1      player.Player
	Player2      player.Player
	Player1Ready bool
	Player2Ready bool
}
