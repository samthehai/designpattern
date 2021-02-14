package builder

import (
	mazeabstract "github.com/samthehai/designpattern/creational/abstractfactory/maze/abstract"
)

type CountingMazeBuilder struct {
	roomCount int
	doorCount int
	wallCount int
}

func (b *CountingMazeBuilder) BuildMaze() {
	b.doorCount = 0
	b.roomCount = 0
	b.wallCount = 0
}

func (b *CountingMazeBuilder) BuildRoom(room int) {
	b.roomCount++
}

func (b *CountingMazeBuilder) BuildDoor(roomFrom, roomTo int) {
	b.roomCount++
}

func (b *CountingMazeBuilder) AddWall(
	room int,
	direction mazeabstract.Direction,
) {
	b.wallCount++
}

func (b *CountingMazeBuilder) GetMaze() mazeabstract.Maze {
	return nil
}

func (b *CountingMazeBuilder) GetCount() (roomCount, doorCount, wallCount int) {
	return b.roomCount, b.doorCount, b.wallCount
}
