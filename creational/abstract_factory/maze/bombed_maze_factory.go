package maze

import "github.com/samthehai/designpattern/creational/abstract_factory/maze/abstract"

type BombedMazeFactory struct {
}

func (m *BombedMazeFactory) MakeMaze() abstract.Maze {
	return &Maze{}
}

func (m *BombedMazeFactory) MakeWall() abstract.Wall {
	return &Wall{}
}

func (m *BombedMazeFactory) MakeRoom(number int) abstract.Room {
	return &Room{
		Number:  number,
		HasBomb: true,
	}
}

func (m *BombedMazeFactory) MakeDoor(room1, room2 abstract.Room) abstract.Door {
	return &Door{
		Room1: room1,
		Room2: room2,
	}
}
