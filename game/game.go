package game

import (
	"github.com/gdamore/tcell"
	"time"
)

const (
	fps              = 2
	canvasBackground = tcell.ColorLightBlue
	canvasPadding    = 10
	paddleHeight     = 5
)

type Game struct {
	ticker      *time.Ticker
	done        chan interface{}
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
		position: Vector{canvasPadding*2 + 2, (termHeight - paddleHeight) / 2},
		height:   paddleHeight,
		color:    tcell.ColorDarkBlue,
	}
	rightPaddle := Paddle{
		position: Vector{termWidth - 2*canvasPadding - 3, (termHeight - paddleHeight) / 2},
		height:   paddleHeight,
		color:    tcell.ColorDarkGreen,
	}

	screen.HideCursor()
	screen.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorOrange))
	screen.Clear()

	return &Game{
		ticker:      time.NewTicker((100 / fps) * time.Millisecond),
		done:        make(chan interface{}),
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
	g.ball.HandleCollision(g)
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

func (g *Game) handlePaddleMove(paddle *Paddle, direction Vector) {
	newPaddleTop := Add(paddle.position, direction).y
	newPaddleBottom := newPaddleTop + paddle.height
	if newPaddleTop >= g.ballCanvas.y && newPaddleBottom <= g.ballCanvas.y+g.ballCanvas.height {
		paddle.Move(direction)
	}
}

func (g *Game) handleKey(ev *tcell.EventKey) {
	switch ev.Rune() {
	case 'W', 'w':
		g.handlePaddleMove(g.leftPaddle, Up())
	case 'S', 's':
		g.handlePaddleMove(g.leftPaddle, Down())

	}
	switch ev.Key() {
	case tcell.KeyCtrlC:
		g.stop()
	case tcell.KeyUp:
		g.handlePaddleMove(g.rightPaddle, Up())
	case tcell.KeyDown:
		g.handlePaddleMove(g.rightPaddle, Down())
	}
}

func (g *Game) stop() {
	g.screen.Fini()
	g.ticker.Stop()
	close(g.done)
}
