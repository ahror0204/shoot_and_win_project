package repository

import (
	"sync"

	"github.com/shoot_and_win/match"
	"github.com/shoot_and_win/player"
)

type InMemoryRepository struct {
	players map[string]player.Player
	matches map[string]match.Match
	mu      *sync.Mutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		players: make(map[string]player.Player),
		matches: make(map[string]match.Match),
		mu:      &sync.Mutex{},
	}
}

func (r *InMemoryRepository) GetMatch(id string) match.Match {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.matches[id]
}

func (r *InMemoryRepository) UpdateMatch(match match.Match) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.matches[match.ID] = match
}

func (r *InMemoryRepository) CreateMatch(match match.Match) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.matches[match.ID] = match
}

func (r *InMemoryRepository) GetPlayer(name string) player.Player {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.players[name]
}

func (r *InMemoryRepository) RemovePlayer(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.players, name)
}

func (r *InMemoryRepository) RegisterPlayer(player player.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.players[player.Name] = player

	return nil
}

func (r *InMemoryRepository) ListPlayers(name string) []player.Player {
	r.mu.Lock()
	defer r.mu.Unlock()

	players := make([]player.Player, 0, len(r.players))
	for key, value := range r.players {
		if key == name {
			continue
		}
		players = append(players, value)
	}

	return players
}

