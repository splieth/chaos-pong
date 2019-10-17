package game

import "github.com/gdamore/tcell"

type Canvas struct {
	x               int
	y               int
	width           int
	height          int
	backgroundColor tcell.Color
}

type Collision int8

const (
	TopWall Collision = iota
	RightWall
	BottomWall
	LeftWall
	LeftPaddle
	RightPaddle
)

func NewCanvas(x, y, width, height int, backgroundColor tcell.Color) Canvas {
	return Canvas{
		x:               x,
		y:               y,
		width:           width,
		height:          height,
		backgroundColor: backgroundColor,
	}
}

func (c *Canvas) Draw(screen tcell.Screen) {
	for col := c.x; col < c.x+c.width; col++ {
		for row := c.y; row < c.y+c.height; row++ {
			screen.SetContent(col, row, ' ', nil, tcell.StyleDefault.Background(c.backgroundColor))
		}
	}
}

func (c *Canvas) GetCenter() Vector {
	return Vector{(c.width / 2) + c.x, (c.height / 2) + c.y}
}
