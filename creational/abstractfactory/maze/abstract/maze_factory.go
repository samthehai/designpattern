package abstract

type MazeFactory interface {
	MakeMaze() Maze
	MakeWall() Wall
	MakeRoom(number int) Room
	MakeDoor(room1, room2 Room) Door
}
