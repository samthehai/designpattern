package maze

import "github.com/samthehai/designpattern/creational/abstract_factory/maze/abstract"

// Maze represents a collection of room
type Maze struct {
	rooms []abstract.Room
}

func (m *Maze) AddRoom(room abstract.Room) {
	if m.rooms == nil {
		m.rooms = make([]abstract.Room, 1)
	}

	m.rooms = append(m.rooms, room)
}

func (m *Maze) RoomNo(number int) abstract.Room {
	if m.rooms == nil {
		return nil
	}

	for _, room := range m.rooms {
		if room != nil && room.GetRoomNumber() == number {
			return room
		}
	}

	return nil
}
