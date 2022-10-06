package main

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) RegisterPlayer(player Player) error {
	err := s.repo.RegisterPlayer(player)
	if err != nil {
		return err
	}
	
	return nil
}