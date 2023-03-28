// define the player model

package models

import (
	"trading-day-go/internal/utils"

	"github.com/gorilla/websocket"
)

// Player represents a player in the game

type Player struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Room *Room `json:"room"`
	Money int `json:"money"`
	Hand []*Card `json:"hand"`

	Socket *websocket.Conn
}

type PlayerRepository interface {
	createPlayer(name string, room *Room) *Player
	findPlayer(id string) *Player
}

type InMemoryPlayerRepository struct {
	players map[string]*Player
}

func createPlayer(r *InMemoryPlayerRepository, name string, room *Room) *Player {
	uuid := utils.CreateUUID()
	player := &Player{
		ID: uuid,
		Name: name,
		Room: room,
	}
	r.players[player.ID] = player
	return player
}