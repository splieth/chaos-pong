package game

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

func (g *Game) handleBallPaddleCollision() {
	ballPos := g.ball.Pos
	leftPaddle := g.player.paddle
	rightPaddle := g.npc.paddle
	if ballPos.X <= leftPaddle.Pos.X+leftPaddle.Width &&
		ballPos.Y <= leftPaddle.Pos.Y+leftPaddle.Height &&
		ballPos.Y >= leftPaddle.Pos.Y {
		g.ball.Dir.InvertX()
		g.ball.Dir.Randomize()
		g.ball.Velocity++
	}
	if ballPos.X+float64(g.ball.Diameter) >= rightPaddle.Pos.X &&
		ballPos.Y <= rightPaddle.Pos.Y+rightPaddle.Height &&
		ballPos.Y >= rightPaddle.Pos.Y {
		g.ball.Dir.InvertX()
		g.ball.Dir.Randomize()
		g.ball.Velocity++
	}
}
