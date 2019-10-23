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

func newBall(canvas *types.Canvas) Ball {
	ballImage := LoadImage("resources/ball.png")
	ballDiameter, _ := ballImage.Size()
	ballRadius := ballDiameter / 2
	return Ball{
		Radius: ballRadius,
		Diameter:ballDiameter,
		Object: types.Object{
			Pos:    canvas.Center,
			Dir:    types.Vector{X: 1, Y: 1},
			Image:  ballImage,
			Canvas: canvas,
		},
	}
}

func (ball *Ball) Move() {
	ball.Pos.Add(ball.Dir)
}

func (ball *Ball) Draw() {
	imageOptions := ebiten.DrawImageOptions{}
	imageOptions.GeoM.Translate(ball.Pos.X, ball.Pos.Y)
	_ = ball.Canvas.Image.DrawImage(ball.Image, &imageOptions)
}
