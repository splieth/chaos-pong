package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/splieth/chaos-pong/game/types"
)

// paddleSpeed is the default movement speed for paddles.
const paddleSpeed = 10

// Player represents a game participant with a side (left/right) and a paddle.
type Player struct {
	side   string
	paddle *Paddle
}

// Paddle is the rectangular controller a player uses to deflect the ball.
type Paddle struct {
	Width, Height float64
	types.Object
}

// NewPlayer creates a player on the given side with a paddle at the
// specified position and color.
func NewPlayer(side string, width, height int, pos types.Vector, clr color.Color, canvas *types.Canvas) Player {
	paddle := NewPaddle(width, height, pos, clr, canvas)
	return Player{
		side:   side,
		paddle: &paddle,
	}
}

// NewPaddle creates a paddle with the given dimensions, position, color,
// and parent canvas.
func NewPaddle(width, height int, pos types.Vector, clr color.Color, canvas *types.Canvas) Paddle {
	image := ebiten.NewImage(width, height)
	image.Fill(clr)
	return Paddle{
		Width:  float64(width),
		Height: float64(height),
		Object: types.Object{
			Pos:      pos,
			Image:    image,
			Canvas:   canvas,
			Velocity: paddleSpeed,
		},
	}
}

// Draw renders the paddle sprite onto the canvas at its current position.
func (p *Player) Draw() {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.paddle.Pos.X, p.paddle.Pos.Y)
	p.paddle.Canvas.Image.DrawImage(p.paddle.Image, opts)
}

// Move applies a directional offset to the paddle, clamped to canvas bounds.
// The offset is normalized and scaled by paddle velocity before being applied.
// A zero-length offset is ignored.
func (p *Player) Move(offset types.Vector) {
	if offset.Norm() == 0 {
		return
	}
	offset.Normalize()
	offset.Multiply(p.paddle.Velocity)
	p.paddle.Pos.Add(offset)
	p.paddle.Pos.Y = max(p.paddle.Pos.Y, 0)
	p.paddle.Pos.Y = min(p.paddle.Pos.Y, p.paddle.Canvas.Height-p.paddle.Height)
}
