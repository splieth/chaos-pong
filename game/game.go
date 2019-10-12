package game

import (
	"github.com/gdamore/tcell"
	"os"
	"time"
)

const (
	fps             = 2
	tick            = (100 / fps) * time.Millisecond
	backgroundColor = tcell.ColorLightBlue
)

func Start(screen tcell.Screen) {
	screen.HideCursor()
	screen.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorOrange))
	screen.Clear()
	loop(screen)
}

// FIXME border detection is worse than before
func (ball *Ball) handleCollision(c *Canvas) {
	newPos := ball.getNextPos()

	if newPos.x > c.x+c.width-CanvasPadding-1 || newPos.x <= c.x {
		ball.sprite.direction.x = ball.sprite.direction.x * -1
	}
	if newPos.y < c.y || newPos.y >= c.y+c.height-CanvasPadding {
		ball.sprite.direction.y = ball.sprite.direction.y * -1
	}
}

func handleKeyPresses(screen tcell.Screen) (bool, tcell.Key) {
	for {
		ev := screen.PollEvent()
		if ev != nil {
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyCtrlQ:
					os.Exit(0)
				}
			}
		}
	}
}

func loop(screen tcell.Screen) {
	go handleKeyPresses(screen)
	termWidth, termHeight := screen.Size()
	ballCanvas := Canvas{
		x:      CanvasPadding,
		y:      CanvasPadding,
		width:  termWidth - CanvasPadding,
		height: termHeight - CanvasPadding,
	}
	ball := Ball{
		sprite: Sprite{
			position:  Vector{(ballCanvas.width / 2) + CanvasPadding, (ballCanvas.height / 2) + CanvasPadding},
			direction: Vector{x: 1, y: 1,},
		},
	}

	t := time.NewTicker(tick)
	for range t.C {
		ball.handleCollision(&ballCanvas)
		ballCanvas.draw(screen)
		ball.move(screen)
		screen.Show()
	}
}
