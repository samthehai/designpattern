package builder

import (
	"github.com/samthehai/designpattern/creational/abstract_factory/maze"
	mazeabstract "github.com/samthehai/designpattern/creational/abstract_factory/maze/abstract"
)

type StandardMazeBuilder struct {
	currentMaze mazeabstract.Maze
}

func (b *StandardMazeBuilder) BuildMaze() {
	b.currentMaze = &maze.Maze{}
}

func (b *StandardMazeBuilder) BuildRoom(room int) {
	if b.currentMaze.RoomNo(room) != nil {
		return
	}

	r := &maze.Room{Number: room}
	b.currentMaze.AddRoom(r)
}

func (b *StandardMazeBuilder) BuildDoor(roomFrom, roomTo int) {
	r1 := b.currentMaze.RoomNo(roomFrom)
	r2 := b.currentMaze.RoomNo(roomTo)
	d := maze.Door{Room1: r1, Room2: r2}

	r1.SetSide(b.commonWall(r1, r2), &d)
	r2.SetSide(b.commonWall(r2, r1), &d)
}

func (b *StandardMazeBuilder) AddWall(
	room int,
	direction mazeabstract.Direction,
) {
	r := b.currentMaze.RoomNo(room)
	if r == nil {
		return
	}

	r.SetSide(direction, &maze.Wall{})
}

func (b *StandardMazeBuilder) GetMaze() mazeabstract.Maze {
	return b.currentMaze
}

func (b *StandardMazeBuilder) commonWall(
	room1, room2 mazeabstract.Room,
) mazeabstract.Direction {
	// TODO: impl
	return mazeabstract.East
}
