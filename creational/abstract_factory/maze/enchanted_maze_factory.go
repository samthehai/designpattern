package maze

import "github.com/samthehai/designpattern/creational/abstract_factory/maze/abstract"

type EnchantedMazeFactory struct {
}

func (m *EnchantedMazeFactory) MakeMaze() abstract.Maze {
	return &Maze{}
}

func (m *EnchantedMazeFactory) MakeWall() abstract.Wall {
	return &Wall{}
}

func (m *EnchantedMazeFactory) MakeRoom(number int) abstract.Room {
	return &Room{
		Number: number,
	}
}

func (m *EnchantedMazeFactory) MakeDoor(room1, room2 abstract.Room) abstract.Door {
	return &Door{
		Room1:       room1,
		Room2:       room2,
		IsEnchanted: true,
	}
}
