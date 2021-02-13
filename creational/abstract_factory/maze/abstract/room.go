package abstract

type Room interface {
	Enter()
	SetSide(direction Direction, mapSite MapSite)
	GetSide(direction Direction) MapSite
	GetRoomNumber() int
}
