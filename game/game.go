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
	ball *types.Ball
}

func NewGame() Game {
	ballImage := LoadImage("resources/ball.png")
	width, _ := ballImage.Size()
	radius := width / 2
	ball := types.NewBall(
		types.Vector{X: 0, Y: 0},
		types.Vector{X: 1, Y: 1},
		ballImage,
		radius)
	return Game{
		ball: &ball,
	}
}

func (game *Game) Draw(screen *ebiten.Image) error {
	game.ball.Draw(screen)
	return nil
}

func LoadImage(path string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return image
}

