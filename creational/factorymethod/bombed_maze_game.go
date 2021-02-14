package factorymethod

import (
	"github.com/samthehai/designpattern/creational/abstractfactory/maze"
	"github.com/samthehai/designpattern/creational/abstractfactory/maze/abstract"
)

type BombedMazeGame struct{}

func (m *BombedMazeGame) MakeMaze() abstract.Maze {
	return &maze.Maze{}
}

func (m *BombedMazeGame) MakeWall() abstract.Wall {
	return &maze.Wall{
		IsEnchanted: false,
	}
}

func (m *BombedMazeGame) MakeRoom(number int) abstract.Room {
	return &maze.Room{
		Number:  number,
		HasBomb: true,
	}
}

func (m *BombedMazeGame) MakeDoor(room1, room2 abstract.Room) abstract.Door {
	return &maze.Door{
		Room1: room1,
		Room2: room2,
	}
}
