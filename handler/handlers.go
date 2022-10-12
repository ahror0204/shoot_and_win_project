package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/shoot_and_win/command"
	"github.com/shoot_and_win/hub"
	"github.com/shoot_and_win/match"
	"github.com/shoot_and_win/player"
	"github.com/shoot_and_win/service"
)


const (
	EventPlayerJoined       = "player_joined"
	EventNewAvailablePlayer = "new_available_player"
	EventMatchCreated       = "match_created"
	EventMatchStarted       = "match_started"
	EventSomeoneWin         = "someone_win"
	EventGameOver           = "game_over"
)

type WebsocketHandler struct {
	s        service.Service
	hub      *hub.Hub
	commands <-chan command.Command
}

func NewWebsocketHandler(s service.Service, h *hub.Hub, comm <-chan command.Command) WebsocketHandler {
	return WebsocketHandler{
		s:        s,
		hub:      h,
		commands: comm,
	}
}

func (h *WebsocketHandler) Run() {
	go func() {
		for cmd := range h.commands {
			fmt.Printf("reacived command %s: %s\n", cmd.Name(), cmd.Payload())
			switch cmd.Name() {
			case command.WaitForOpponent:
				var payload WaitForOpponentPayload
				err := json.Unmarshal(cmd.Payload(), &payload)
				if err != nil {
					log.Printf("failed to unmarshal %s payload: %v\n", cmd.Name(), err)
					continue
				}
				h.s.WaitForSomeone(payload.Player.Name)
				event := Event{
					Name:   EventNewAvailablePlayer,
					Player: payload.Player,
				}

				h.notifyOthers(payload.Player.Name, event.Marshal())

			case command.Play:
				var payload PlayPayload
				err := json.Unmarshal(cmd.Payload(), &payload)
				if err != nil {
					log.Printf("failed to unmarshal %s payload: %v\n", cmd.Name(), err)
					continue
				}

				player := h.s.GetPlayer(payload.Player.Name)
				rival := h.s.GetPlayer(payload.Rival.Name)
				event := Event{
					Name:   EventPlayerJoined,
					Player: payload.Player,
				}

				err = h.hub.Write(rival, event.Marshal())
				if err != nil {
					log.Printf("failed to write event to %s: %v\n", rival.Name, err)
					continue
				}
				id := uuid.NewString()

				m := match.Match{
					ID: id,
					Player1: player,
					Player2: rival,
					Player1Ready: false,
					Player2Ready: false,
				}
				h.s.CreateMatch(m)
				
				event = Event{
					Name: EventMatchCreated,
					MetaData: map[string]interface{}{
						"match_id": m.ID,
					},
				}

				h.hub.Write(player, event.Marshal())
				h.hub.Write(rival, event.Marshal())

			case command.Start:
				var payload StartPayload
				err := json.Unmarshal(cmd.Payload(), &payload)
				if err != nil {
					log.Printf("failed to unmarshal %s payload: %v\n", cmd.Name(), err)
					continue
				}

				m := h.s.GetMatch(payload.MatchID)
				if m.Player1.Name == payload.Player.Name {
					m.Player1Ready = true
				}
				if m.Player2.Name == payload.Player.Name {
					m.Player2Ready = true
				}

				h.s.UpdateMatch(m)

				if m.Player1Ready && m.Player2Ready {
					event := Event{
						Name: EventMatchStarted,
						MetaData: map[string]interface{}{
							"match_id": m.ID,
						},
					}
					h.hub.Write(m.Player1, event.Marshal())
					h.hub.Write(m.Player2, event.Marshal())
				}
			case command.Shoot:
				
			default:
				panic("no such command")
			}
		}
	}()
}

func (h *WebsocketHandler) notifyOthers(self string, event []byte) {
	players := h.s.AllPlayers()
	for _, p := range players {
		if p.Name == self {
			continue
		}

		err := h.hub.Write(p, event)
		if err != nil {
			log.Printf("failed to write in notifyAll: %v\n", err)
			continue
		}
	}
}

type StartPayload struct{
	MatchID string `json:"match_id"`
	Player player.Player `json:"player"`
}

type WaitForOpponentPayload struct {
	Player player.Player `json:"player"`
}

type PlayPayload struct {
	Player player.Player `json:"player"`
	Rival  player.Player `json:"rival"`
}

type Event struct {
	Name     string                 `json:"name"`
	Player   player.Player          `json:"player"`
	MetaData map[string]interface{} `json:"meta_data"`
}

func (e Event) Marshal() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return data
}
