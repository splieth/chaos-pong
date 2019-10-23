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
			vel:    direction,
			image:  image,
			canvas: canvas,
		},
	}
}

func (b *Ball) Move() {
	canvasWidth, canvasHeight := b.canvas.Image.Size()
	if b.pos.X < 0 {
		b.vel.X = -b.vel.X
	}
	if b.pos.X > float64(canvasWidth-2*b.Radius) {
		b.vel.X = -b.vel.X
	}
	if b.pos.Y < 0 {
		b.vel.Y = -b.vel.Y
	}
	if b.pos.Y > float64(canvasHeight-2*b.Radius) {
		b.vel.Y = -b.vel.Y
	}

	b.pos.X = b.pos.X + b.vel.X
	b.pos.Y = b.pos.Y + b.vel.Y
}

func (ball *Ball) Draw() {
	imageOptions := ebiten.DrawImageOptions{}
	imageOptions.GeoM.Translate(ball.pos.X, ball.pos.Y)
	_ = ball.canvas.Image.DrawImage(ball.image, &imageOptions)
}
