package maze

import "github.com/samthehai/designpattern/creational/abstractfactory/maze/abstract"

type NormalMazeFactory struct {
}

func (m *NormalMazeFactory) MakeMaze() abstract.Maze {
	return &Maze{}
}

func (m *NormalMazeFactory) MakeWall() abstract.Wall {
	return &Wall{}
}

func (m *NormalMazeFactory) MakeRoom(number int) abstract.Room {
	return &Room{
		Number: number,
	}
}

func (m *NormalMazeFactory) MakeDoor(room1, room2 abstract.Room) abstract.Door {
	return &Door{
		Room1: room1,
		Room2: room2,
	}
}
