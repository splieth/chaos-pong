package game

// Wall represents which wall the ball has collided with.
type Wall int

const (
	TopWall Wall = iota
	RightWall
	BottomWall
	LeftWall
	NoWall
)

// handleBallCanvasCollision checks if the ball has hit any canvas boundary.
// On collision, the ball's direction is inverted along the appropriate axis.
// Returns which wall was hit, or NoWall if no collision occurred.
func (g *Game) handleBallCanvasCollision() Wall {
	bounds := g.ball.Canvas.Image.Bounds()
	canvasWidth, canvasHeight := bounds.Dx(), bounds.Dy()

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

// handleBallPaddleCollision checks if the ball overlaps with either paddle.
// On hit, the ball's horizontal direction is inverted, a random vertical
// perturbation is applied, and the ball speed increases.
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
