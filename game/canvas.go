package game

import "github.com/gdamore/tcell"

const (
	CanvasPadding = 4
)

type Canvas struct {
	x      int
	y      int
	width  int
	height int
}

func (c *Canvas) draw(screen tcell.Screen) {
	for col := c.x; col < c.width; col++ {
		for row := c.y; row < c.height; row++ {
			screen.SetContent(col, row, ' ', nil, tcell.StyleDefault.Background(backgroundColor))
		}
	}
}
