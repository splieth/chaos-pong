package game

import (
	"math"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/splieth/chaos-pong/game/types"
	"github.com/stretchr/testify/assert"
)

func newTestBall(pos, dir types.Vector, velocity float64) Ball {
	canvas := types.NewCanvas(types.Vector{X: 0, Y: 0}, 800, 600)
	image := ebiten.NewImage(10, 10)
	return Ball{
		Radius:   5,
		Diameter: 10,
		Object: types.Object{
			Pos:      pos,
			Dir:      dir,
			Velocity: velocity,
			Image:    image,
			Canvas:   &canvas,
		},
	}
}

func TestMove_AdvancesPosition(t *testing.T) {
	// given
	ball := newTestBall(
		types.Vector{X: 100, Y: 100},
		types.Vector{X: 1, Y: 0},
		10,
	)

	// when
	ball.Move()

	// then
	assert.InDelta(t, 110, ball.Pos.X, 1e-10)
	assert.InDelta(t, 100, ball.Pos.Y, 1e-10)
}

func TestMove_DiagonalDirection(t *testing.T) {
	// given
	ball := newTestBall(
		types.Vector{X: 100, Y: 100},
		types.Vector{X: 1, Y: 1},
		10,
	)

	// when
	ball.Move()

	// then
	expectedDelta := 10.0 / math.Sqrt(2)
	assert.InDelta(t, 100+expectedDelta, ball.Pos.X, 1e-10)
	assert.InDelta(t, 100+expectedDelta, ball.Pos.Y, 1e-10)
}

func TestMove_DirectionRemainsNormalized(t *testing.T) {
	// given
	ball := newTestBall(
		types.Vector{X: 100, Y: 100},
		types.Vector{X: 3, Y: 4},
		5,
	)

	// when
	ball.Move()

	// then
	norm := ball.Dir.Norm()
	assert.InDelta(t, 1.0, norm, 1e-10)
}

func TestMove_RespectsVelocity(t *testing.T) {
	// given
	slow := newTestBall(types.Vector{X: 0, Y: 0}, types.Vector{X: 1, Y: 0}, 5)
	fast := newTestBall(types.Vector{X: 0, Y: 0}, types.Vector{X: 1, Y: 0}, 20)

	// when
	slow.Move()
	fast.Move()

	// then
	assert.InDelta(t, 5.0, slow.Pos.X, 1e-10)
	assert.InDelta(t, 20.0, fast.Pos.X, 1e-10)
}

func TestMove_MultipleTicks(t *testing.T) {
	// given
	ball := newTestBall(
		types.Vector{X: 0, Y: 0},
		types.Vector{X: 1, Y: 0},
		10,
	)

	// when
	ball.Move()
	ball.Move()
	ball.Move()

	// then
	assert.InDelta(t, 30.0, ball.Pos.X, 1e-10)
	assert.InDelta(t, 0.0, ball.Pos.Y, 1e-10)
}

func TestDraw_DoesNotPanic(t *testing.T) {
	// given
	ball := newTestBall(
		types.Vector{X: 50, Y: 50},
		types.Vector{X: 1, Y: 0},
		10,
	)

	// when / then
	assert.NotPanics(t, func() {
		ball.Draw()
	})
}
