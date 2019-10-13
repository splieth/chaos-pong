package game

import "github.com/gdamore/tcell"

type Canvas struct {
	x      int
	y      int
	width  int
	height int
}

func (c *Canvas) Draw(screen tcell.Screen) {
	for col := c.x; col < c.x+c.width; col++ {
		for row := c.y; row < c.y+c.height; row++ {
			screen.SetContent(col, row, ' ', nil, tcell.StyleDefault.Background(canvasBackground))
		}
	}
}
