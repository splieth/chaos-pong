package game

import (
	"github.com/gdamore/tcell"
)

type Ball struct {
	position  IntVector
	direction FloatVector
	color     tcell.Color
}

func NewBall(pos IntVector, dir FloatVector, color tcell.Color) Ball {
	return Ball{
		position:  pos,
		direction: dir,
		color:     color,
	}
}

func (ball *Ball) GetNextPos() IntVector {
	return ball.position.convertToFloat().Add(ball.direction).convertToInt()
}

func (ball *Ball) Move() {
	ball.position = ball.GetNextPos()
}

func (ball *Ball) Draw(screen tcell.Screen) {
	roundedBallPos := ball.position
	screen.SetContent(roundedBallPos.x, roundedBallPos.y, '●', nil,
		tcell.StyleDefault.Background(ballCanvasBG).Foreground(ball.color))
}
