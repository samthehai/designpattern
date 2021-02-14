package factorymethod

import (
	"github.com/samthehai/designpattern/creational/abstractfactory/maze"
	"github.com/samthehai/designpattern/creational/abstractfactory/maze/abstract"
)

type EnchantedMazeGame struct{}

func (m *EnchantedMazeGame) MakeMaze() abstract.Maze {
	return &maze.Maze{}
}

func (m *EnchantedMazeGame) MakeWall() abstract.Wall {
	return &maze.Wall{}
}

func (m *EnchantedMazeGame) MakeRoom(number int) abstract.Room {
	return &maze.Room{
		Number: number,
	}
}

func (m *EnchantedMazeGame) MakeDoor(room1, room2 abstract.Room) abstract.Door {
	return &maze.Door{
		Room1:       room1,
		Room2:       room2,
		IsEnchanted: true,
	}
}
