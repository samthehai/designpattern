package factorymethod_test

import (
	"testing"

	"github.com/samthehai/designpattern/creational/factorymethod"
	"github.com/stretchr/testify/assert"
)

func TestCreateMaze(t *testing.T) {
	type TestCase struct {
		Name     string
		GameType factorymethod.MazeGameType
		HasError bool
	}

	tests := []TestCase{
		{
			Name:     "bombed maze game",
			GameType: factorymethod.MazeGameTypeBombed,
		},
		{
			Name:     "enchanted maze game",
			GameType: factorymethod.MazeGameTypeBombed,
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			game, err := factorymethod.GetMazeGame(tc.GameType)
			if tc.HasError {
				assert.Error(t, err)
				assert.Nil(t, game)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, game)
			}
		})
	}
}
