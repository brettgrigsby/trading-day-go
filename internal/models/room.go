// define the Room model

package models

import (
	"trading-day-go/internal/utils"
)

type Room struct {
	ID string `json:"id"`
	Players []*Player `json:"players"`
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