package maze

import "log"

type Wall struct {
	IsEnchanted bool
	HasBomb     bool
}

func (w *Wall) Enter() {
	log.Println("Enter wall")
}
