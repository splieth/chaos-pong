package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/splieth/chaos-pong/game"
	"log"
	"os"
)

func main() {
	basePath := os.Getenv("PWD")
	pong := game.NewGame(1280, 720, basePath)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Chaos Pong!")

	if err := ebiten.RunGame(&pong); err != nil {
		log.Fatal(err)
	}
}
