package types

import (
	"github.com/hajimehoshi/ebiten"
)

type Ball struct {
	Radius int
	Object
}

func NewBall(startPosition, direction Vector, canvas *Canvas, image *ebiten.Image, radius int) Ball {
	return Ball{
		Radius: radius,
		Object: Object{
			pos:    startPosition,
			dir:    direction,
			image:  image,
			canvas: canvas,
		},
	}
}

func (b *Ball) Move() {
	canvasWidth, canvasHeight := b.canvas.Image.Size()
	if b.pos.X < 0 {
		b.dir.X = -b.dir.X
	}
	if b.pos.X > float64(canvasWidth-2*b.Radius) {
		b.dir.X = -b.dir.X
	}
	if b.pos.Y < 0 {
		b.dir.Y = -b.dir.Y
	}
	if b.pos.Y > float64(canvasHeight-2*b.Radius) {
		b.dir.Y = -b.dir.Y
	}

	b.pos.X = b.pos.X + b.dir.X
	b.pos.Y = b.pos.Y + b.dir.Y
}

func (ball *Ball) Draw() {
	imageOptions := ebiten.DrawImageOptions{}
	imageOptions.GeoM.Translate(ball.pos.X, ball.pos.Y)
	_ = ball.canvas.Image.DrawImage(ball.image, &imageOptions)
}
