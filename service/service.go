package service

import (
	"fmt"

	"github.com/shoot_and_win/match"
	"github.com/shoot_and_win/player"
)

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

type Service struct {
	repo Repository
}


func (s Service) RemoveMatch(id string) {
	s.repo.RemoveMatch(id)
}

func (s Service) GetMatch(matchID string) match.Match{
	return s.repo.GetMatch(matchID)
}

func (s Service) UpdateMatch(match match.Match) {
	s.repo.CreateMatch(match)
}

func (s Service) CreateMatch(match match.Match) {
	s.repo.UpdateMatch(match)
}

func (s Service) GetPlayer(name string) player.Player {
	return s.repo.GetPlayer(name)
}

func (s Service) RemovePlayer(name string) {
	s.repo.RemovePlayer(name)
}

func (s Service) SavePlayer(player player.Player) ([]player.Player, error) {
	err := s.repo.SavePlayer(player)
	if err != nil {
		return nil, err
	}

	listOfPlayers := s.repo.ListPlayers()

	return listOfPlayers, nil
}

func (s Service) CreatePlayer(p player.Player){
	p.Health = 100
	s.repo.SavePlayer(p)
}

func (s Service) WaitForSomeone(name string) {
	p := s.GetPlayer(name)
	p.SetWaitingForOpponent(true)
	s.repo.SavePlayer(p)
}

func (s Service) AvailablePlayers() []player.Player {
	availablePlayers := make([]player.Player, 0)
	players := s.repo.ListPlayers()
	fmt.Println("Available players: ", players)
	for _, p := range players {
		if p.IsWaitingForOpponent() {
			availablePlayers = append(availablePlayers, p)
		}
	}
	return availablePlayers
}

func (s Service) AllPlayers() []player.Player {
	return s.repo.ListPlayers()
}
