package main

type Repository interface{
	RegisterPlayer(player Player) error
	ListPlayers(name string) []Player
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
	r.players[player.Name] = player

	return nil
}


func (r *InMemoryRepository) ListPlayers(name string) []Player {
	players := make([]Player, 0, len(r.players))
	
	for key, value := range r.players {
		if key == name {
			continue
		}
		players = append(players, value)
	}

	return players
}
