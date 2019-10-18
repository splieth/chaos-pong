package types

import (
	"github.com/hajimehoshi/ebiten"
)

type Ball struct {
	Radius int
	Object
}

func NewBall(startPosition, direction Vector, image *ebiten.Image, radius int) Ball {
	return Ball{
		Radius: radius,
		Object: Object{
			pos:   startPosition,
			vel:   direction,
			image: image,
		},
	}
}

func (ball *Ball) Move(screen *ebiten.Image) {
	imageOptions := ebiten.DrawImageOptions{}

	screenWidth, screenHeight := screen.Size()
	if ball.pos.X < 0 {
		ball.vel.X = -ball.vel.X
	}
	if ball.pos.X > float64(screenWidth-ball.Radius) {
		ball.vel.X = -ball.vel.X
	}
	if ball.pos.Y < 0 {
		ball.vel.Y = -ball.vel.Y
	}
	if ball.pos.Y > float64(screenHeight-ball.Radius) {
		ball.vel.Y = -ball.vel.Y
	}

	ball.pos.X = ball.pos.X + ball.vel.X
	ball.pos.Y = ball.pos.Y + ball.vel.Y
	imageOptions.GeoM.Translate(ball.pos.X, ball.pos.Y)
	_ = screen.DrawImage(ball.image, &imageOptions)
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	ball.Move(screen)
}
