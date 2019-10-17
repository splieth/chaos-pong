package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_randomize(t *testing.T) {
	t.Run("random addition is in range", func(t *testing.T) {
		input := IntVector{
			x: 0,
			y: 0,
		}
		firstRandomAddition := randomize(input)
		secondRandomAddition := randomize(input)

		assert.Less(t, firstRandomAddition, 1.0)
		assert.Greater(t, firstRandomAddition, 0.0)
		assert.NotEqual(t, firstRandomAddition, secondRandomAddition)
	})
}
