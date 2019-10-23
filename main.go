package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/splieth/chaos-pong/game"
	"log"
)

var (
	pong game.Game
)

func main() {
	screen, _ := ebiten.NewImage(500, 500, ebiten.FilterDefault)
	width, height := screen.Size()
	pong = game.NewGame(screen)

	if err := ebiten.Run(pong.Draw, width, height, 1, "Chaos Pong!"); err != nil {
		log.Fatal(err)
	}
}
