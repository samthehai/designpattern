package builder_test

import (
	"testing"

	"github.com/samthehai/designpattern/creational/builder"
	"github.com/samthehai/designpattern/creational/builder/abstract"
	"github.com/stretchr/testify/assert"
)

func TestCreateMaze(t *testing.T) {
	type TestCase struct {
		Name          string
		Builder       abstract.MazeBuilder
		IsOnlyCouting bool
	}

	tests := []TestCase{
		{
			Name:    "standard maze builder",
			Builder: &builder.StandardMazeBuilder{},
		},
		{
			Name:          "counting maze builder",
			Builder:       &builder.CountingMazeBuilder{},
			IsOnlyCouting: true,
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			game := &builder.MazeGame{}
			game.CreateMaze(tc.Builder)

			maze := tc.Builder.GetMaze()
			if tc.IsOnlyCouting {
				assert.Nil(t, maze)
			} else {
				assert.NotNil(t, maze)
			}
		})
	}
}
