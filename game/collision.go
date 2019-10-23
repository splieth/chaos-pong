package game

import (
	"github.com/splieth/chaos-pong/game/types"
)

func (g *Game) handleBallCanvasCollision() {
	canvasWidth, canvasHeight := g.ball.Canvas.Image.Size()
	if g.ball.Pos.X < 0 || g.ball.Pos.X > float64(canvasWidth-g.ball.Diameter) {
		g.ball.Dir.InvertX()
	}
	if g.ball.Pos.Y < 0 || g.ball.Pos.Y > float64(canvasHeight-g.ball.Diameter) {
		g.ball.Dir.InvertY()
	}
}

func (g *Game) handlePaddelCanvasCollision() {
	down := types.Vector{X: 0, Y: 1}
	up := types.Vector{X: 0, Y: -1}
	if g.leftPaddle.Pos.Y < 0 {
		g.leftPaddle.Pos.Add(down)
	}
	if g.leftPaddle.Pos.Y+g.leftPaddle.Height > g.ballCanvas.Height {
		g.leftPaddle.Pos.Add(up)
	}
	if g.rightPaddle.Pos.Y < 0 {
		g.rightPaddle.Pos.Add(down)
	}
	if g.rightPaddle.Pos.Y+g.rightPaddle.Height > g.ballCanvas.Height {
		g.rightPaddle.Pos.Add(up)
	}
}
