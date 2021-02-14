package factorymethod

import (
	"errors"
	"fmt"

	"github.com/samthehai/designpattern/creational/abstractfactory/maze/abstract"
)

type MazeGame interface {
	MakeMaze() abstract.Maze
	MakeRoom(room int) abstract.Room
	MakeDoor(roomFrom, roomTo abstract.Room) abstract.Door
	MakeWall() abstract.Wall
}

type MazeGameType int

const (
	MazeGameTypeBombed MazeGameType = iota
	MazeGameTypeEnchanted
)

func GetMazeGame(g MazeGameType) (MazeGame, error) {
	switch g {
	case MazeGameTypeBombed:
		return new(BombedMazeGame), nil
	case MazeGameTypeEnchanted:
		return new(EnchantedMazeGame), nil
	default:
		return nil, errors.New(fmt.Sprintf("Maze game %d was not recognized", g))
	}
}
