package game

import (
	"fmt"
	"github.com/gdamore/tcell"
	"strings"
	"time"
)

const (
	initialFps    = 2
	ballCanvasBG  = tcell.ColorLightBlue
	scoreCanvasBG = tcell.ColorBlack
	canvasPadding = 10
	paddleHeight  = 5
	goalSleepTime = 500 * time.Millisecond
)

type Game struct {
	ticker      *time.Ticker
	done        chan interface{}
	screen      tcell.Screen
	ballCanvas  *Canvas
	scoreCanvas *Canvas
	ball        *Ball
	leftPaddle  *Paddle
	rightPaddle *Paddle
	scores      []int
	fps         int
}

func NewGame(screen tcell.Screen) *Game {
	termWidth, termHeight := screen.Size()
	ballCanvas := NewCanvas(2*canvasPadding, canvasPadding, termWidth-4*canvasPadding, termHeight-2*canvasPadding, ballCanvasBG)
	scoreCanvas := NewCanvas(2*canvasPadding, termHeight-canvasPadding, termWidth-4*canvasPadding, 1, scoreCanvasBG)
	ball := NewBall(ballCanvas.GetCenter(), FloatVector{1, 1}, tcell.ColorOrangeRed)
	leftPaddle := NewPaddle(IntVector{canvasPadding*2 + 2, (termHeight - paddleHeight) / 2}, paddleHeight, tcell.ColorDarkBlue)
	rightPaddle := NewPaddle(IntVector{termWidth - 2*canvasPadding - 3, (termHeight - paddleHeight) / 2}, paddleHeight, tcell.ColorDarkGreen)

	screen.HideCursor()
	screen.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorOrange))
	screen.Clear()

	return &Game{
		ticker:      frameRate(initialFps),
		done:        make(chan interface{}),
		screen:      screen,
		ballCanvas:  &ballCanvas,
		scoreCanvas: &scoreCanvas,
		ball:        &ball,
		leftPaddle:  &leftPaddle,
		rightPaddle: &rightPaddle,
		scores:      []int{0, 0},
	}
}

func (g *Game) EventLoop() {
	defer g.ticker.Stop()
	go g.pollScreenEvents()
	go g.increaseBallSpeed()

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
	g.HandleBallCollision()
	g.move()
	g.draw()
}

func (g *Game) stop() {
	g.screen.Fini()
	g.ticker.Stop()
	close(g.done)
}

func (g *Game) scoreGoal(collision Collision) {
	if collision == RightWall {
		g.scores[0]++
	} else {
		g.scores[1]++
	}
	g.ball.position = g.ballCanvas.GetCenter()
	g.resetBallSpeed()
	g.ball.direction.x = g.ball.direction.x * -1
	g.ball.direction.y = g.ball.direction.y * -1
	time.Sleep(goalSleepTime)
}

func (g *Game) move() {
	g.ball.Move()
}

func (g *Game) draw() {
	g.ballCanvas.Draw(g.screen)
	g.scoreCanvas.Draw(g.screen)
	g.ball.Draw(g.screen)
	g.leftPaddle.Draw(g.screen)
	g.rightPaddle.Draw(g.screen)
	g.updateScores()
	g.screen.Show()
}

func (g *Game) updateScores() {
	scoreString := strings.Trim(strings.Replace(fmt.Sprint(g.scores), " ", ":", -1), "[]")
	textAnchor := (g.scoreCanvas.width / 2) + canvasPadding + len(scoreString)
	for i, r := range scoreString {
		g.screen.SetContent(textAnchor+i, g.scoreCanvas.y, r, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
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

func (g *Game) handlePaddleMove(paddle *Paddle, direction IntVector) {
	newPaddleTop := paddle.position.Add(direction).y
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

func frameRate(fps int) *time.Ticker {
	return time.NewTicker(time.Duration((100 / fps)) * time.Millisecond)
}

func (g *Game) increaseBallSpeed() {
	g.fps++
	g.ticker = frameRate(g.fps)
}
func (g *Game) resetBallSpeed() {
	g.fps = initialFps
	g.ticker = frameRate(initialFps)
}
