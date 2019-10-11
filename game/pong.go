package game

import (
	"github.com/gdamore/tcell"
	"os"
	"time"
)

const FPS = 2
const GAME_HEIGHT = 32
const GAME_WIDTH = 64

type Coordinate struct {
	x int
	y int
}

type Direction struct {
	xStep int
	yStep int
}

type Ball struct {
	position Coordinate
	dir      Direction
	speed    int
}

func (ball *Ball) getNextPos() Coordinate {
	newX := ball.position.x + ball.dir.xStep*ball.speed
	newY := ball.position.y + ball.dir.yStep*ball.speed
	return Coordinate{newX, newY}
}

func (ball *Ball) move() {
	ball.position = ball.getNextPos()
}

func (ball *Ball) handleCollision(topCorner Coordinate, bottomCorner Coordinate) {
	newPos := ball.getNextPos()
	if newPos.x <= topCorner.x || newPos.x >= bottomCorner.x {
		ball.dir.xStep = ball.dir.xStep * -1
	}
	if newPos.y <= topCorner.y || newPos.y >= bottomCorner.y {
		ball.dir.yStep = ball.dir.yStep * -1
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
				case tcell.KeyESC:
					os.Exit(0)
				}
			}
		}
	}
}

func drawSquare(screen tcell.Screen, coord Coordinate, color tcell.Color) {
	screen.SetCell(coord.x, coord.y, tcell.StyleDefault.Background(color), ' ')
	screen.SetCell(coord.x+1, coord.y, tcell.StyleDefault.Background(color), ' ')
}

func drawRectangle(screen tcell.Screen, topCorner Coordinate, bootomCorner Coordinate) {
	for x := topCorner.x; x <= bootomCorner.x; x++ {
		drawSquare(screen, Coordinate{x, topCorner.y}, tcell.ColorWhite)
		drawSquare(screen, Coordinate{x, bootomCorner.y}, tcell.ColorWhite)
	}
	for y := topCorner.y; y <= bootomCorner.y; y++ {
		drawSquare(screen, Coordinate{topCorner.x, y}, tcell.ColorWhite)
		drawSquare(screen, Coordinate{bootomCorner.x, y}, tcell.ColorWhite)
	}
}

func draw(screen tcell.Screen, topCorner Coordinate, bottomCorner Coordinate, ball Ball) {
	screen.Clear()
	drawRectangle(screen, topCorner, bottomCorner)
	screen.SetCell(ball.position.x, ball.position.y, tcell.StyleDefault, '●')
	screen.Show()
}

func loop(screen tcell.Screen) {
	go handleKeyPresses(screen)
	termWidth, termHeight := screen.Size()
	topCorner := Coordinate{termWidth/2 - GAME_WIDTH, termHeight/2 - GAME_HEIGHT/2}
	bottomCorner := Coordinate{termWidth/2 + GAME_WIDTH, termHeight/2 + GAME_HEIGHT/2}
	ball := Ball{Coordinate{termWidth / 2, termHeight / 2}, Direction{1, 1}, 1}
	for {
		ball.handleCollision(topCorner, bottomCorner)
		ball.move()
		draw(screen, topCorner, bottomCorner, ball)
		duration := (100 / FPS) * time.Millisecond
		time.Sleep(duration)
	}
}
