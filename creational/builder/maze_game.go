package builder

import (
	mazeabstract "github.com/samthehai/designpattern/creational/abstractfactory/maze/abstract"
	"github.com/samthehai/designpattern/creational/builder/abstract"
)

type MazeGame struct{}

func (g *MazeGame) CreateMaze(builder abstract.MazeBuilder) mazeabstract.Maze {
	builder.BuildMaze()
	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}
