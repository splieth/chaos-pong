package game

import (
	"github.com/gdamore/tcell"
	"time"
)

const (
	fps             = 2
	backgroundColor = tcell.ColorLightBlue
)

type Game struct {
	ticker     *time.Ticker
	done       chan bool
	screen     tcell.Screen
	ballCanvas *Canvas
	ball       *Ball
}

func NewGame(screen tcell.Screen) *Game {
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

	screen.HideCursor()
	screen.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorOrange))
	screen.Clear()

	return &Game{
		ticker:     time.NewTicker((100 / fps) * time.Millisecond),
		done:       make(chan bool),
		screen:     screen,
		ballCanvas: &ballCanvas,
		ball:       &ball,
	}
}

func (g *Game) EventLoop() {
	defer g.ticker.Stop()
	go g.pollScreenEvents()

	for {
		select {
		case <-g.done:
			return
		case <-g.ticker.C:
			g.tick()
		}
	}
}

func (g *Game) tick() {
	g.ball.handleCollision(g.ballCanvas)
	g.move()
	g.draw()
}

func (g *Game) move() {
	g.ball.move()
}

func (g *Game) draw() {
	g.ballCanvas.draw(g.screen)
	g.ball.draw(g.screen)
	g.screen.Show()
}

func (g *Game) pollScreenEvents() {
	for {
		ev := g.screen.PollEvent()
		if ev != nil {
			switch ev := ev.(type) {
			case *tcell.EventKey:
				g.handleKey(ev)
			}
		}
	}
}

func (g *Game) handleKey(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyCtrlQ:
		g.stop()
	}
}

func (g *Game) stop() {
	g.screen.Fini()
	g.ticker.Stop()
	g.done <- true
}

// FIXME border detection is worse than before
// FIXME this needs to go somewhere else
func (ball *Ball) handleCollision(c *Canvas) {
	newPos := ball.getNextPos()

	if newPos.x > c.x+c.width-CanvasPadding-1 || newPos.x <= c.x {
		ball.sprite.direction.x = ball.sprite.direction.x * -1
	}
	if newPos.y < c.y || newPos.y >= c.y+c.height-CanvasPadding {
		ball.sprite.direction.y = ball.sprite.direction.y * -1
	}
}
