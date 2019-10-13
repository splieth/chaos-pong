package game

import "github.com/gdamore/tcell"

type Ball struct {
	position  Vector
	direction Vector
	color     tcell.Color
}

func (ball *Ball) Move() {
	ball.position = Add(ball.position, ball.direction)
}

func (ball *Ball) Draw(screen tcell.Screen) {
	screen.SetContent(ball.position.x, ball.position.y, '●', nil, tcell.StyleDefault.Background(canvasBackground).Foreground(ball.color))
}

func (ball *Ball) HandleCollision(c *Canvas) {
	newPos := Add(ball.position, ball.direction)

	if newPos.x >= c.x+c.width || newPos.x < c.x {
		ball.direction.x = ball.direction.x * -1
	}
	if newPos.y < c.y || newPos.y >= c.y+c.height {
		ball.direction.y = ball.direction.y * -1
	}
}
