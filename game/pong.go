package game

import (
	"github.com/gdamore/tcell"
	"os"
	"time"
)

const (
	FPS           = 3
	CanvasPadding = 4
)

type Vector struct {
	x, y int
}

type Sprite struct {
	position  Vector
	direction Vector
}

type Ball struct {
	sprite Sprite
	speed  int
}

type Canvas struct {
	x      int
	y      int
	width  int
	height int
}

func (c *Canvas) draw(screen tcell.Screen) {
	for col := c.x; col < c.width; col++ {
		for row := c.y; row < c.height; row++ {
			screen.SetContent(col, row, ' ', nil, tcell.StyleDefault.Background(tcell.ColorRebeccaPurple))
		}
	}
}

func Start(screen tcell.Screen) {
	screen.HideCursor()
	screen.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorOrange))
	screen.Clear()
	loop(screen)
}

func (ball *Ball) getNextPos() Vector {
	newX := ball.sprite.position.x + ball.sprite.direction.x*ball.speed
	newY := ball.sprite.position.y + ball.sprite.direction.y*ball.speed
	return Vector{newX, newY}
}

func (ball *Ball) move(screen tcell.Screen) {
	ball.sprite.position = ball.getNextPos()
	screen.SetContent(ball.sprite.position.x, ball.sprite.position.y, '●', nil, tcell.StyleDefault.Background(tcell.ColorRebeccaPurple).Foreground(tcell.ColorOrangeRed))
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
		speed: 1,
	}
	for {
		ball.handleCollision(&ballCanvas)
		ballCanvas.draw(screen)
		ball.move(screen)
		screen.Show()
		duration := (100 / FPS) * time.Millisecond
		time.Sleep(duration)
	}
}
