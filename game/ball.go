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
	ball.Pos = ball.Pos.Add(ball.Dir)
}

func (ball *Ball) Draw() {
	imageOptions := ebiten.DrawImageOptions{}
	imageOptions.GeoM.Translate(ball.Pos.X, ball.Pos.Y)
	_ = ball.Canvas.Image.DrawImage(ball.Image, &imageOptions)
}
