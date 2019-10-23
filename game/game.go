package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/splieth/chaos-pong/game/types"
	"image/color"
	"log"
)

const (
	canvasPadding = 50.0
	paddleWidth   = 25
	paddleHeight  = 150
)

type Game struct {
	ball        *Ball
	ballCanvas  *types.Canvas
	leftPaddle  *Paddle
	rightPaddle *Paddle
}

func NewGame(screen *ebiten.Image) Game {
	canvas := types.NewCanvas(screen, types.Vector{X: canvasPadding, Y: canvasPadding}, canvasPadding)
	canvasWidth, _ := canvas.Image.Size()

	ball := newBall(&canvas)

	leftPaddleColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	rightPaddleColor := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	leftPaddlePos := types.Vector{X: 0, Y: 0}
	rightPaddlePos := types.Vector{X: float64(canvasWidth - paddleWidth), Y: 0}

	leftPaddle := NewPaddle(paddleWidth, paddleHeight, leftPaddlePos, leftPaddleColor, &canvas)
	rightPaddle := NewPaddle(paddleWidth, paddleHeight, rightPaddlePos, rightPaddleColor, &canvas)

	return Game{
		ball:        &ball,
		ballCanvas:  &canvas,
		leftPaddle:  &leftPaddle,
		rightPaddle: &rightPaddle,
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ballCanvas.Fill()
	g.ball.Draw()
	g.leftPaddle.Draw()
	g.rightPaddle.Draw()
	g.ballCanvas.Draw(screen)
}

func (g *Game) Tick(screen *ebiten.Image) error {
	leftPaddleOffset, rightPaddleOffset := getPaddleMoves()
	g.handleBallCanvasCollision()
	g.ball.Move()
	g.handlePaddelCanvasCollision()
	g.leftPaddle.Move(leftPaddleOffset)
	g.rightPaddle.Move(rightPaddleOffset)
	g.Draw(screen)
	return nil
}

func LoadImage(path string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return image
}
