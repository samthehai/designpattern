package maze

import (
	"log"

	"github.com/samthehai/designpattern/creational/abstractfactory/maze/abstract"
)

type Room struct {
	HasBomb  bool
	Number   int
	mapSites [4]abstract.MapSite
}

func (r *Room) Enter() {
	log.Println("Enter a normal room")
}

func (r *Room) SetSide(direction abstract.Direction, mapSite abstract.MapSite) {
	r.mapSites[int(direction)] = mapSite
}

func (r *Room) GetSide(direction abstract.Direction) abstract.MapSite {
	return r.mapSites[int(direction)]
}

func (r *Room) GetRoomNumber() int {
	return r.Number
}
