package game

import (
	"github.com/hajimehoshi/ebiten"
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
	ballImage := createBallImage()

	ball := types.NewBall(
		types.Vector{X: 0, Y: 0},
		types.Vector{X: 1, Y: 1},
		ballImage,
		BallRadius)
	return Game{
		ball: &ball,
	}
}

func (game *Game) Draw(screen *ebiten.Image) error {
	game.ball.Draw(screen)
	return nil
}

func createBallImage() *ebiten.Image {
	image, err := ebiten.NewImage(2*BallRadius, 2*BallRadius, ebiten.FilterNearest)
	if err != nil {
		log.Panic(err)
	}
	return image
}
