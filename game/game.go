package game

import (
	"github.com/gdamore/tcell"
	"time"
)

const (
	fps              = 2
	canvasBackground = tcell.ColorLightBlue
	canvasPadding    = 10
)

type Game struct {
	ticker      *time.Ticker
	done        chan bool
	screen      tcell.Screen
	ballCanvas  *Canvas
	ball        *Ball
	leftPaddle  *Paddle
	rightPaddle *Paddle
}

func NewGame(screen tcell.Screen) *Game {
	termWidth, termHeight := screen.Size()
	ballCanvas := Canvas{
		x:      2 * canvasPadding,
		y:      canvasPadding,
		width:  termWidth - 4*canvasPadding,
		height: termHeight - 2*canvasPadding,
	}
	ball := Ball{
		position:  Vector{(ballCanvas.width / 2) + canvasPadding, (ballCanvas.height / 2) + canvasPadding},
		direction: Vector{1, 1},
		color:     tcell.ColorOrangeRed,
	}
	leftPaddle := Paddle{
		position: Vector{canvasPadding * 2, canvasPadding},
		height:   5,
		color:    tcell.ColorDarkBlue,
	}
	rightPaddle := Paddle{
		position: Vector{termWidth - 2*canvasPadding - 1, canvasPadding},
		height:   5,
		color:    tcell.ColorDarkGreen,
	}

	screen.HideCursor()
	screen.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorOrange))
	screen.Clear()

	return &Game{
		ticker:      time.NewTicker((100 / fps) * time.Millisecond),
		done:        make(chan bool),
		screen:      screen,
		ballCanvas:  &ballCanvas,
		ball:        &ball,
		leftPaddle:  &leftPaddle,
		rightPaddle: &rightPaddle,
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
	g.ball.HandleCollision(g.ballCanvas)
	g.move()
	g.draw()
}

func (g *Game) move() {
	g.ball.Move()
}

func (g *Game) draw() {
	g.ballCanvas.Draw(g.screen)
	g.ball.Draw(g.screen)
	g.leftPaddle.Draw(g.screen)
	g.rightPaddle.Draw(g.screen)
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
	switch ev.Rune() {
	case 'W', 'w':
		if Add(g.leftPaddle.position, Up()).y >= g.ballCanvas.y {
			g.leftPaddle.Move(Up())
		}
	case 'S', 's':
		if Add(g.leftPaddle.position, Down()).y+g.leftPaddle.height <= g.ballCanvas.y+g.ballCanvas.height {
			g.leftPaddle.Move(Down())
		}

	}
	switch ev.Key() {
	case tcell.KeyCtrlC:
		g.stop()
	case tcell.KeyUp:
		if Add(g.rightPaddle.position, Up()).y >= g.ballCanvas.y {
			g.rightPaddle.Move(Up())
		}
	case tcell.KeyDown:
		if Add(g.rightPaddle.position, Down()).y+g.rightPaddle.height <= g.ballCanvas.y+g.ballCanvas.height {
			g.rightPaddle.Move(Down())
		}
	}
}

func (g *Game) stop() {
	g.screen.Fini()
	g.ticker.Stop()
	g.done <- true
}
