package game

func (g *Game) handleBallCanvasCollision() {
	canvasWidth, canvasHeight := g.ball.Canvas.Image.Size()
	if g.ball.Pos.X < 0 {
		g.ball.Dir.X = -g.ball.Dir.X
	}
	if g.ball.Pos.X > float64(canvasWidth-g.ball.Diameter) {
		g.ball.Dir.X = -g.ball.Dir.X
	}
	if g.ball.Pos.Y < 0 {
		g.ball.Dir.Y = -g.ball.Dir.Y
	}
	if g.ball.Pos.Y > float64(canvasHeight-g.ball.Diameter) {
		g.ball.Dir.Y = -g.ball.Dir.Y
	}
}
