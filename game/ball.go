package game

import (
	"github.com/gdamore/tcell"
)

type Ball struct {
	position  Vector
	direction Vector
	color     tcell.Color
}

func NewBall(pos, dir Vector, color tcell.Color) Ball {
	return Ball{
		position:  pos,
		direction: dir,
		color:     color,
	}
}

func (ball *Ball) Move() {
	ball.position = Add(ball.position, ball.direction)
}

func (ball *Ball) Draw(screen tcell.Screen) {
	screen.SetContent(ball.position.x, ball.position.y, '●', nil,
		tcell.StyleDefault.Background(ballCanvasBG).Foreground(ball.color))
}
