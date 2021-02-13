package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/samthehai/designpattern/creational/abstract_factory/maze"
	"github.com/samthehai/designpattern/creational/abstract_factory/maze/abstract"
)

func main() {
	spew.Dump(createMaze(&maze.NormalMazeFactory{}))
	spew.Dump(createMaze(&maze.BombedMazeFactory{}))
	spew.Dump(createMaze(&maze.EnchantedMazeFactory{}))
}

func createMaze(factory abstract.MazeFactory) abstract.Maze {
	maze := factory.MakeMaze()
	r1 := factory.MakeRoom(1)
	r2 := factory.MakeRoom(2)
	door := factory.MakeDoor(r1, r2)

	maze.AddRoom(r1)
	maze.AddRoom(r2)

	r1.SetSide(abstract.North, factory.MakeWall())
	r1.SetSide(abstract.East, door)
	r1.SetSide(abstract.South, factory.MakeWall())
	r1.SetSide(abstract.West, factory.MakeWall())

	r2.SetSide(abstract.North, factory.MakeWall())
	r2.SetSide(abstract.East, factory.MakeWall())
	r2.SetSide(abstract.South, factory.MakeWall())
	r2.SetSide(abstract.West, door)

	return maze
}
