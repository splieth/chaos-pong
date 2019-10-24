package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/splieth/chaos-pong/game/types"
)

type Ball struct {
	Radius   int
	Diameter int
	types.Object
}

func newBall(canvas *types.Canvas, basePath string) Ball {
	ballImage := LoadImage(basePath, "/resources/ball.png")
	ballDiameter, _ := ballImage.Size()
	ballRadius := ballDiameter / 2
	return Ball{
		Radius:   ballRadius,
		Diameter: ballDiameter,
		Object: types.Object{
			Pos:      canvas.Center,
			Dir:      types.Vector{X: 1, Y: 1},
			Image:    ballImage,
			Canvas:   canvas,
			Velocity: 10,
		},
	}
}

func (b *Ball) Move() {
	b.Dir.Normalize()
	b.Dir.Multiply(b.Velocity)
	b.Pos.Add(b.Dir)
}

func (b *Ball) Draw() {
	imageOptions := ebiten.DrawImageOptions{}
	imageOptions.GeoM.Translate(b.Pos.X, b.Pos.Y)
	_ = b.Canvas.Image.DrawImage(b.Image, &imageOptions)
}
