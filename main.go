package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/splieth/chaos-pong/game"
	"log"
	"os"
)

func initScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("Could not start tcell for chaos-pong.")
		log.Printf("Cannot alloc screen, tcell.NewScreen() gave an error:\n%s", err)
		os.Exit(1)
	}
	err = screen.Init()
	if err != nil {
		fmt.Println("Could not start tcell for chaos-pong.")
		log.Printf("Cannot start gomatrix, screen.Init() gave an error:\n%s", err)
		os.Exit(1)
	}
	return screen
}

func main() {
	screen := initScreen()
	game.Start(screen)
}
