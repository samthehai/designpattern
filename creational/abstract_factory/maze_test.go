package maze_test

import (
	"testing"

	"github.com/samthehai/designpattern/creational/abstract_factory/maze"
	"github.com/samthehai/designpattern/creational/abstract_factory/maze/abstract"
	"github.com/stretchr/testify/assert"
)

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

func TestCreateMaze(t *testing.T) {
	type TestCase struct {
		Name        string
		MazeFactory abstract.MazeFactory
	}

	tests := []TestCase{
		{
			Name:        "normal maze factory",
			MazeFactory: &maze.NormalMazeFactory{},
		},
		{
			Name:        "bombed maze factory",
			MazeFactory: &maze.BombedMazeFactory{},
		},
		{
			Name:        "enchanted maze factory",
			MazeFactory: &maze.EnchantedMazeFactory{},
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			maze := createMaze(tc.MazeFactory)
			assert.NotNil(t, maze)
		})
	}
}
