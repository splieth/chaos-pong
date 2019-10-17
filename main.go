package main

import (
	"github.com/gdamore/tcell"
	"github.com/splieth/chaos-pong/game"
	"log"
	"os"
	"time"
)

func initScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Printf("Cannot alloc screen, tcell.NewScreen() gave an error:\n%s", err)
		os.Exit(1)
	}
	if err = screen.Init(); err != nil {
		log.Printf("Cannot start gomatrix, screen.Init() gave an error:\n%s", err)
		os.Exit(1)
	}
	return screen
}

func setupLogging() {
	t := time.Now()
	tmsmp := (t.Format("20060102150405"))
	f, err := os.OpenFile("testlogfile-"+tmsmp+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
}

func main() {
	setupLogging()
	screen := initScreen()
	g := game.NewGame(screen)
	g.EventLoop()
}
