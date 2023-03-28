// define the Room model

package models

import (
	"trading-day-go/internal/utils"
)

type Room struct {
	ID string `json:"id"`
	Players []*Player `json:"players"`
	DrawDeck []*Card `json:"drawDeck"`
	DiscardDeck []*Card `json:"discardDeck"`
	CurrentPlayer *Player `json:"currentPlayer"`
	EventCardDrawDeck []*EventCard `json:"eventCardDrawDeck"`
	EventCardOnDeck []*EventCard `json:"eventCardOnDeck"`
	EventCardHistory []*EventCard `json:"eventCardHistory"`
}

type RoomRepository interface {
	createRoom() *Room
	findRoom(id string) *Room
}

type InMemoryRoomRepository struct {
	rooms map[string]*Room
}

func (r *InMemoryRoomRepository) createRoom() *Room {
	uuid := utils.CreateUUID()
	room := &Room{
		ID: uuid,
		Players: []*Player{},
	}
	r.rooms[room.ID] = room
	return room
}

func findRoom(r *InMemoryRoomRepository, id string) *Room {
	return r.rooms[id]
}