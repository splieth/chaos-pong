package game

import (
	"github.com/gdamore/tcell"
	"os"
	"time"
)

const (
	FPS        = 2
	GameHeight = 32
	GameWidth  = 64
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

func (ball *Ball) getNextPos() Vector {
	newX := ball.sprite.position.x + ball.sprite.direction.x*ball.speed
	newY := ball.sprite.position.y + ball.sprite.direction.y*ball.speed
	return Vector{newX, newY}
}

func (ball *Ball) move() {
	ball.sprite.position = ball.getNextPos()
}

func (ball *Ball) handleCollision(topCorner Vector, bottomCorner Vector) {
	newPos := ball.getNextPos()
	if newPos.x <= topCorner.x || newPos.x >= bottomCorner.x {
		ball.sprite.direction.x = ball.sprite.direction.x * -1
	}
	if newPos.y <= topCorner.y || newPos.y >= bottomCorner.y {
		ball.sprite.direction.y = ball.sprite.direction.y * -1
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

func drawSquare(screen tcell.Screen, coord Vector, color tcell.Color) {
	screen.SetCell(coord.x, coord.y, tcell.StyleDefault.Background(color), ' ')
	screen.SetCell(coord.x+1, coord.y, tcell.StyleDefault.Background(color), ' ')
}

func drawRectangle(screen tcell.Screen, topCorner Vector, bootomCorner Vector) {
	for x := topCorner.x; x <= bootomCorner.x; x++ {
		drawSquare(screen, Vector{x, topCorner.y}, tcell.ColorWhite)
		drawSquare(screen, Vector{x, bootomCorner.y}, tcell.ColorWhite)
	}
	for y := topCorner.y; y <= bootomCorner.y; y++ {
		drawSquare(screen, Vector{topCorner.x, y}, tcell.ColorWhite)
		drawSquare(screen, Vector{bootomCorner.x, y}, tcell.ColorWhite)
	}
}

func draw(screen tcell.Screen, topCorner Vector, bottomCorner Vector, ball Ball) {
	screen.Clear()
	drawRectangle(screen, topCorner, bottomCorner)
	screen.SetCell(ball.sprite.position.x, ball.sprite.position.y, tcell.StyleDefault, '●')
	screen.Show()
}

func loop(screen tcell.Screen) {
	go handleKeyPresses(screen)
	termWidth, termHeight := screen.Size()
	topCorner := Vector{termWidth/2 - GameWidth, termHeight/2 - GameHeight/2}
	bottomCorner := Vector{termWidth/2 + GameWidth, termHeight/2 + GameHeight/2}
	ball := Ball{
		sprite: Sprite{
			position:  Vector{termWidth / 2, termHeight / 2},
			direction: Vector{x: 1, y: 1,},
		},
		speed: 1,
	}
	for {
		ball.handleCollision(topCorner, bottomCorner)
		ball.move()
		draw(screen, topCorner, bottomCorner, ball)
		duration := (100 / FPS) * time.Millisecond
		time.Sleep(duration)
	}
}
