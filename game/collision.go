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
	leftPaddlePos := g.leftPaddle.Pos
	rightPaddlePos := g.rightPaddle.Pos
	if ballPos.X <= leftPaddlePos.X+g.leftPaddle.Width &&
		ballPos.Y <= leftPaddlePos.Y+g.leftPaddle.Height &&
		ballPos.Y >= leftPaddlePos.Y {
		g.ball.Dir.InvertX()
		g.ball.Dir.Randomize()
		g.ball.Velocity++
	}
	if ballPos.X+float64(g.ball.Diameter) >= rightPaddlePos.X &&
		ballPos.Y <= rightPaddlePos.Y+g.rightPaddle.Height &&
		ballPos.Y >= rightPaddlePos.Y {
		g.ball.Dir.InvertX()
		g.ball.Dir.Randomize()
		g.ball.Velocity++
	}
}
