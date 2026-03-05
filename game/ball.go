package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/splieth/chaos-pong/game/types"
)

// InitialBallSpeed is the starting velocity of the ball at the beginning
// of each round. Speed increases with each paddle hit.
const InitialBallSpeed = 10

// Ball represents the game ball with collision dimensions derived from its sprite.
type Ball struct {
	Radius   int
	Diameter int
	types.Object
}

// newBall creates a ball centered on the given canvas, loading its sprite
// from the resources directory.
func newBall(canvas *types.Canvas, basePath string) Ball {
	ballImage := LoadImage(basePath, "/resources/ball.png")
	diameter := ballImage.Bounds().Dx()
	return Ball{
		Radius:   diameter / 2,
		Diameter: diameter,
		Object: types.Object{
			Pos:      canvas.Center,
			Dir:      types.Vector{X: 1, Y: 1},
			Image:    ballImage,
			Canvas:   canvas,
			Velocity: InitialBallSpeed,
		},
	}
}

// Move advances the ball position by one tick. The direction vector is
// normalized, scaled by velocity, applied to position, then re-normalized
// so velocity changes don't affect direction.
func (b *Ball) Move() {
	b.Dir.Normalize()
	b.Dir.Multiply(b.Velocity)
	b.Pos.Add(b.Dir)
	b.Dir.Normalize()
}

// Draw renders the ball sprite onto its canvas at the current position.
func (b *Ball) Draw() {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(b.Pos.X, b.Pos.Y)
	b.Canvas.Image.DrawImage(b.Image, opts)
}
