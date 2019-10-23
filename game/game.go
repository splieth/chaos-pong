package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/splieth/chaos-pong/game/types"
	"image/color"
	"log"
)

type Game struct {
	ball        *types.Ball
	ballCanvas  *types.Canvas
	leftPaddle  *types.Paddle
	rightPaddle *types.Paddle
}

func getCenter(screen *ebiten.Image) (float64, float64) {
	width, height := screen.Size()
	return float64(width) / 2, float64(height) / 2
}

func NewGame(screen *ebiten.Image) Game {
	ballImage := LoadImage("resources/ball.png")
	ballWidth, _ := ballImage.Size()
	radius := ballWidth / 2
	posX, posY := getCenter(screen)
	saize := 250
	image, _ := ebiten.NewImage(saize, saize, ebiten.FilterDefault)
	leftPaddleImage, _ := ebiten.NewImage(50, 100, ebiten.FilterDefault)
	leftPaddleImage.Fill(color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})

	rightPaddleImage, _ := ebiten.NewImage(50, 100, ebiten.FilterDefault)
	rightPaddleImage.Fill(color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	})

	canvas := types.Canvas{
		X:     posX - float64(saize/2),
		Y:     posY - float64(saize/2),
		Color: color.White,
		Image: image,
	}

	ball := types.NewBall(
		types.Vector{X: float64(saize / 2), Y: float64(saize / 2)},
		types.Vector{X: 1, Y: 1},
		&canvas,
		ballImage,
		radius)

	leftPaddle := types.NewPaddle(
		types.Vector{X: 0, Y: 0},
		types.Vector{X: 0, Y: 0},
		&canvas,
		leftPaddleImage)

	rightPaddle := types.NewPaddle(
		types.Vector{X: float64(saize - 50), Y: float64(saize - 150)},
		types.Vector{X: 0, Y: 0},
		&canvas,
		rightPaddleImage)

	return Game{
		ball:        &ball,
		ballCanvas:  &canvas,
		leftPaddle:  &leftPaddle,
		rightPaddle: &rightPaddle,
	}
}

func (g *Game) drawCanvas(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	width, height := screen.Size()
	options.GeoM.Translate(float64(width/4), float64(height/4))
	screen.DrawImage(g.ballCanvas.Image, &options)
}

func getPaddleMoves() (types.Vector, types.Vector) {
	leftPaddleOffset := types.Vector{0, 0}
	rightPaddleOffset := types.Vector{0, 0}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		leftPaddleOffset = types.Vector{0, -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		leftPaddleOffset = types.Vector{0, 1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		rightPaddleOffset = types.Vector{0, -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		rightPaddleOffset = types.Vector{0, 1}
	}
	return leftPaddleOffset, rightPaddleOffset
}

func (g *Game) Draw(screen *ebiten.Image) error {
	g.ball.Move()
	leftPaddleOffset, rightPaddleOffset := getPaddleMoves()
	g.leftPaddle.Move(leftPaddleOffset)
	g.rightPaddle.Move(rightPaddleOffset)
	g.ballCanvas.Fill()
	g.ball.Draw()
	g.leftPaddle.Draw()
	g.rightPaddle.Draw()
	g.drawCanvas(screen)
	return nil
}

func LoadImage(path string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return image
}
