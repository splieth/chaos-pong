package game

import (
	"github.com/splieth/chaos-pong/game/types"
)

var (
	paddleUp   = types.Vector{X: 0, Y: 1}
	paddleDown = types.Vector{X: 0, Y: -1}
)

type Wall int

const (
	TopWall = iota
	RightWall
	BottomWall
	LeftWall
	NoWall
)

func (g *Game) handleBallCanvasCollision() Wall {
	canvasWidth, canvasHeight := g.ball.Canvas.Image.Size()
	if g.ball.Pos.Y < 0 {
		g.ball.Dir.InvertY()
		return TopWall
	}
	if g.ball.Pos.X > float64(canvasWidth-g.ball.Diameter) {
		g.ball.Dir.InvertX()
		return RightWall
	}
	if g.ball.Pos.Y > float64(canvasHeight-g.ball.Diameter) {
		g.ball.Dir.InvertY()
		return BottomWall
	}
	if g.ball.Pos.X < 0 {
		g.ball.Dir.InvertX()
		return LeftWall
	}
	return NoWall
}

func (g *Game) handlePaddleCanvasCollision() {
	if g.leftPaddle.Pos.Y < 0 {
		g.leftPaddle.Pos.Add(paddleUp)
	}
	if g.leftPaddle.Pos.Y+g.leftPaddle.Height >= g.ballCanvas.Height {
		g.leftPaddle.Pos.Add(paddleDown)
	}
	if g.rightPaddle.Pos.Y < 0 {
		g.rightPaddle.Pos.Add(paddleUp)
	}
	if g.rightPaddle.Pos.Y+g.rightPaddle.Height >= g.ballCanvas.Height {
		g.rightPaddle.Pos.Add(paddleDown)
	}
}

func (g *Game) handleBallPaddleCollision() {
	ballPos := g.ball.Pos
	leftPaddlePos := g.leftPaddle.Pos
	rightPaddlePos := g.rightPaddle.Pos
	if ballPos.X <= leftPaddlePos.X+g.leftPaddle.Width &&
		ballPos.Y <= leftPaddlePos.Y+g.leftPaddle.Height &&
		ballPos.Y >= leftPaddlePos.Y {
		g.ball.Dir.InvertX()
		g.ball.Dir.Randomize()
	}
	if ballPos.X+float64(g.ball.Diameter) >= rightPaddlePos.X &&
		ballPos.Y <= rightPaddlePos.Y+g.rightPaddle.Height &&
		ballPos.Y >= rightPaddlePos.Y {
		g.ball.Dir.InvertX()
		g.ball.Dir.Randomize()
	}
}
