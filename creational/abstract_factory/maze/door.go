package maze

import (
	"log"

	"github.com/samthehai/designpattern/creational/abstract_factory/maze/abstract"
)

type Door struct {
	Room1       abstract.Room
	Room2       abstract.Room
	IsOpen      bool
	IsEnchanted bool
}

func (d *Door) Enter() {
	log.Println("Enter door")
}
