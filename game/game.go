package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/splieth/chaos-pong/game/types"
	"image/color"
	"log"
)

type Game struct {
	ball       *types.Ball
	ballCanvas *types.Canvas
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
	return Game{
		ball:       &ball,
		ballCanvas: &canvas,
	}
}

func (g *Game) drawCanvas(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	width, height := screen.Size()
	options.GeoM.Translate(float64(width/4), float64(height/4))
	screen.DrawImage(g.ballCanvas.Image, &options)
}

func (g *Game) Draw(screen *ebiten.Image) error {
	g.ball.Move()
	g.ballCanvas.Fill()
	g.ball.Draw()
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
