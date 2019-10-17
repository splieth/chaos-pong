package game

import (
	"github.com/gdamore/tcell"
)

type Ball struct {
	position          FloatVector
	lastDrawnPosition IntVector
	direction         FloatVector
	color             tcell.Color
}

func NewBall(pos FloatVector, dir FloatVector, color tcell.Color) Ball {
	return Ball{
		position:  pos,
		direction: dir,
		color:     color,
	}
}

func (ball *Ball) GetNextPos() FloatVector {
	normalizedBallDir := ball.direction.Normalize()
	return ball.position.Add(normalizedBallDir)
}

func (ball *Ball) Move() {
	ball.position = ball.GetNextPos()
}

func (ball *Ball) Draw(screen tcell.Screen) {
	lastPos := ball.lastDrawnPosition
	currPos := ball.position.convertToInt()
	nextPos := ball.GetNextPos().convertToInt()
	if (lastPos.x != nextPos.x && lastPos.y != nextPos.y) &&
		(lastPos.x == currPos.x || lastPos.y == currPos.y) &&
		(nextPos.x == currPos.x || nextPos.y == currPos.y) {
		screen.SetContent(lastPos.x, lastPos.y, '●', nil,
			tcell.StyleDefault.Background(ballCanvasBG).Foreground(ball.color))
	} else {
		screen.SetContent(currPos.x, currPos.y, '●', nil,
			tcell.StyleDefault.Background(ballCanvasBG).Foreground(ball.color))
		ball.lastDrawnPosition = ball.position.convertToInt()
	}
}
