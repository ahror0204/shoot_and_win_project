package main

type Repository interface{
	RegisterPlayer(player Player) error
}

type InMemoryRepository struct {
	players map[string]Player
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		players: make(map[string]Player),
	}
}

func (r *InMemoryRepository) RegisterPlayer(player Player) error {
	r.players[player.ID] = player

	return nil
}


