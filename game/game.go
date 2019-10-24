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
	player1       = "player1"
	player2       = "player2"
)

type Game struct {
	ball        *Ball
	ballCanvas  *types.Canvas
	leftPaddle  *Paddle
	rightPaddle *Paddle
	score       map[string]int
}

func NewGame(screen *ebiten.Image, basePath string) Game {
	canvas := types.NewCanvas(screen, types.Vector{X: canvasPadding, Y: canvasPadding}, canvasPadding)
	canvasWidth, _ := canvas.Image.Size()

	ball := newBall(&canvas, basePath)

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
		score: map[string]int{
			player1: 0,
			player2: 0,
		},
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ballCanvas.Fill()
	g.ball.Draw()
	g.leftPaddle.Draw()
	g.rightPaddle.Draw()
	g.ballCanvas.Draw(screen)
}

func (g *Game) handleScores(wall Wall) {
	if wall == RightWall {
		g.score[player1]++
		log.Println(g.score)
	}
	if wall == LeftWall {
		g.score[player2]++
		log.Println(g.score)
	}
}

func (g *Game) Tick(screen *ebiten.Image) error {
	handleExit()
	collidedWall := g.handleBallCanvasCollision()
	leftPaddleOffset, rightPaddleOffset := getPaddleMoves()
	g.leftPaddle.Move(leftPaddleOffset)
	g.rightPaddle.Move(rightPaddleOffset)
	g.handleScores(collidedWall)
	g.handleBallPaddleCollision()
	g.ball.Move()
	g.Draw(screen)
	return nil
}

func LoadImage(resourcesBasePath, path string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(resourcesBasePath+path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return image
}
