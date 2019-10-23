package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/splieth/chaos-pong/game/types"
)

type Ball struct {
	Radius int
	types.Object
}

func newBall(canvas *types.Canvas) Ball {
	ballImage := LoadImage("resources/ball.png")
	ballDiameter, _ := ballImage.Size()
	ballRadius := ballDiameter / 2
	canvasWidth, canvasHeight := canvas.Image.Size()
	return Ball{
		Radius: ballRadius,
		Object: types.Object{
			Pos:    types.Vector{X: float64(canvasWidth / 2), Y: float64(canvasHeight / 2)},
			Dir:    types.Vector{X: 1, Y: 1},
			Image:  ballImage,
			Canvas: canvas,
		},
	}
}

func (ball *Ball) Move() {
	canvasWidth, canvasHeight := ball.Canvas.Image.Size()
	if ball.Pos.X < 0 {
		ball.Dir.X = -ball.Dir.X
	}
	if ball.Pos.X > float64(canvasWidth-2*ball.Radius) {
		ball.Dir.X = -ball.Dir.X
	}
	if ball.Pos.Y < 0 {
		ball.Dir.Y = -ball.Dir.Y
	}
	if ball.Pos.Y > float64(canvasHeight-2*ball.Radius) {
		ball.Dir.Y = -ball.Dir.Y
	}

	ball.Pos.X = ball.Pos.X + ball.Dir.X
	ball.Pos.Y = ball.Pos.Y + ball.Dir.Y
}

func (ball *Ball) Draw() {
	imageOptions := ebiten.DrawImageOptions{}
	imageOptions.GeoM.Translate(ball.Pos.X, ball.Pos.Y)
	_ = ball.Canvas.Image.DrawImage(ball.Image, &imageOptions)
}
