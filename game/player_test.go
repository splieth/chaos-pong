package game

import (
	"image/color"
	"testing"

	"github.com/splieth/chaos-pong/game/types"
	"github.com/stretchr/testify/assert"
)

func newTestPlayer(side string, pos types.Vector, canvasWidth, canvasHeight float64) *Player {
	canvas := types.NewCanvas(types.Vector{X: 0, Y: 0}, canvasWidth, canvasHeight)
	p := NewPlayer(side, paddleWidth, paddleHeight, pos, color.RGBA{R: 255, A: 255}, &canvas)
	return &p
}

func TestNewPlayer_Side(t *testing.T) {
	// given / when
	player := newTestPlayer("left", types.Vector{X: 0, Y: 0}, 800, 600)

	// then
	assert.Equal(t, "left", player.side)
}

func TestNewPlayer_PaddleDimensions(t *testing.T) {
	// given / when
	player := newTestPlayer("left", types.Vector{X: 0, Y: 0}, 800, 600)

	// then
	assert.Equal(t, float64(paddleWidth), player.paddle.Width)
	assert.Equal(t, float64(paddleHeight), player.paddle.Height)
}

func TestNewPlayer_PaddlePosition(t *testing.T) {
	// given
	pos := types.Vector{X: 100, Y: 200}

	// when
	player := newTestPlayer("right", pos, 800, 600)

	// then
	assert.Equal(t, pos, player.paddle.Pos)
}

func TestNewPlayer_PaddleVelocity(t *testing.T) {
	// given / when
	player := newTestPlayer("left", types.Vector{X: 0, Y: 0}, 800, 600)

	// then
	assert.Equal(t, float64(paddleSpeed), player.paddle.Velocity)
}

func TestNewPlayer_PaddleImage(t *testing.T) {
	// given / when
	player := newTestPlayer("left", types.Vector{X: 0, Y: 0}, 800, 600)

	// then
	assert.NotNil(t, player.paddle.Image)
	w, h := player.paddle.Image.Bounds().Dx(), player.paddle.Image.Bounds().Dy()
	assert.Equal(t, paddleWidth, w)
	assert.Equal(t, paddleHeight, h)
}

func TestPlayerMove_MovesDown(t *testing.T) {
	// given
	player := newTestPlayer("left", types.Vector{X: 0, Y: 100}, 800, 600)

	// when
	player.Move(types.Vector{X: 0, Y: 1})

	// then
	assert.InDelta(t, 100+paddleSpeed, player.paddle.Pos.Y, 1e-10)
}

func TestPlayerMove_MovesUp(t *testing.T) {
	// given
	player := newTestPlayer("left", types.Vector{X: 0, Y: 100}, 800, 600)

	// when
	player.Move(types.Vector{X: 0, Y: -1})

	// then
	assert.InDelta(t, 100-paddleSpeed, player.paddle.Pos.Y, 1e-10)
}

func TestPlayerMove_ClampedAtTop(t *testing.T) {
	// given
	player := newTestPlayer("left", types.Vector{X: 0, Y: 5}, 800, 600)

	// when
	player.Move(types.Vector{X: 0, Y: -1})

	// then
	assert.Equal(t, 0.0, player.paddle.Pos.Y)
}

func TestPlayerMove_ClampedAtBottom(t *testing.T) {
	// given
	bottomY := 600.0 - float64(paddleHeight)
	player := newTestPlayer("left", types.Vector{X: 0, Y: bottomY - 3}, 800, 600)

	// when
	player.Move(types.Vector{X: 0, Y: 1})

	// then
	assert.Equal(t, bottomY, player.paddle.Pos.Y)
}

func TestPlayerMove_ZeroOffsetIgnored(t *testing.T) {
	// given
	player := newTestPlayer("left", types.Vector{X: 0, Y: 200}, 800, 600)

	// when
	player.Move(types.Vector{X: 0, Y: 0})

	// then
	assert.Equal(t, 200.0, player.paddle.Pos.Y)
}

func TestPlayerMove_XPositionUnchanged(t *testing.T) {
	// given
	player := newTestPlayer("left", types.Vector{X: 50, Y: 200}, 800, 600)

	// when
	player.Move(types.Vector{X: 0, Y: 1})

	// then
	assert.Equal(t, 50.0, player.paddle.Pos.X)
}

func TestPlayerDraw_DoesNotPanic(t *testing.T) {
	// given
	player := newTestPlayer("left", types.Vector{X: 0, Y: 100}, 800, 600)

	// when / then
	assert.NotPanics(t, func() {
		player.Draw()
	})
}
