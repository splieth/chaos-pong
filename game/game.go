package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/splieth/chaos-pong/game/types"
	"log"
)

const (
	BallRadius = 15
)

type Game struct {
	ball       *types.Ball
	ballCanvas *ebiten.Image
}

func NewGame(screen *ebiten.Image) Game {
	ballImage := LoadImage("resources/ball.png")
	ballWidth, _ := ballImage.Size()
	radius := ballWidth / 2
	canvas, _ := ebiten.NewImage(250, 250, ebiten.FilterDefault)

	ball := types.NewBall(
		types.Vector{X: 0, Y: 0},
		types.Vector{X: 1, Y: 1},
		canvas,
		ballImage,
		radius)
	return Game{
		ball:       &ball,
		ballCanvas: canvas,
	}
}
func (g *Game) fillCanvas() {
	img := LoadImage("resources/grass.png")
	options := ebiten.DrawImageOptions{}
	g.ballCanvas.DrawImage(img, &options)
}
func (g *Game) drawCanvas(screen *ebiten.Image) {
	screen.DrawImage(g.ballCanvas, &ebiten.DrawImageOptions{})
}

func (g *Game) Draw(screen *ebiten.Image) error {
	g.ball.Move()
	g.fillCanvas()
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
