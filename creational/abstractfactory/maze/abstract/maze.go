package abstract

type Maze interface {
	AddRoom(room Room)
	RoomNo(number int) Room
}
