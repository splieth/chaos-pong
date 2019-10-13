package game

import "github.com/gdamore/tcell"

type Paddle struct {
	position Vector
	color    tcell.Color
	height   int
}

func (paddle *Paddle) Draw(screen tcell.Screen) {
	for y := paddle.position.y; y < paddle.position.y+paddle.height; y++ {
		screen.SetContent(paddle.position.x, y, ' ', nil, tcell.StyleDefault.Background(paddle.color))
	}
}

func (paddle *Paddle) Move(offset Vector) {
	paddle.position = Add(paddle.position, offset)
}
