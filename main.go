package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/splieth/chaos-pong/game"
	"log"
)

var (
	pong game.Game
)

func init() {
	pong = game.NewGame()
}

func main() {
	screen, _ := ebiten.NewImage(320, 240, ebiten.FilterDefault)
	width, height := screen.Size()
	if err := ebiten.Run(pong.Draw, width, height, 1, "Chaos Pong!"); err != nil {
		log.Fatal(err)
	}
}
