package service

import (
	"github.com/shoot_and_win/match"
	"github.com/shoot_and_win/player"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) UpdateMatch(match match.Match) {
	s.CreateMatch(match)
}

func (s Service) CreateMatch(match match.Match) {
	s.UpdateMatch(match)
}

func (s Service) GetPlayer(name string) player.Player {
	return s.GetPlayer(name)
}

func (s Service) RemovePlayer(name string) {
	s.RemovePlayer(name)
}

func (s Service) RegisterPlayer(player player.Player) ([]player.Player, error) {
	err := s.repo.RegisterPlayer(player)
	if err != nil {
		return nil, err
	}

	listOfPlayers := s.repo.ListPlayers(player.Name)

	return listOfPlayers, nil
}
