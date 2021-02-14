package abstract

import (
	mazeabstract "github.com/samthehai/designpattern/creational/abstract_factory/maze/abstract"
)

type MazeBuilder interface {
	BuildMaze()
	BuildRoom(room int)
	BuildDoor(roomFrom, roomTo int)
	AddWall(room int, direction mazeabstract.Direction)
	GetMaze() mazeabstract.Maze
}
