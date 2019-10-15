package game

import "github.com/gdamore/tcell"

type Paddle struct {
	position      Vector
	lastDirection Vector
	color         tcell.Color
	height        int
}

func NewPaddle(position Vector, height int, color tcell.Color) Paddle {
	return Paddle{
		position:      position,
		lastDirection: Vector{0, 0},
		color:         color,
		height:        height,
	}
}

func (paddle *Paddle) Draw(screen tcell.Screen) {
	for y := paddle.position.y; y < paddle.position.y+paddle.height; y++ {
		screen.SetContent(paddle.position.x, y, ' ', nil, tcell.StyleDefault.Background(paddle.color))
	}
}

func (paddle *Paddle) Move(offset Vector) {
	paddle.lastDirection = offset
	paddle.position = Add(paddle.position, offset)
}
