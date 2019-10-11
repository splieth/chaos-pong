package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"log"
	"os"
)

func main() {
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
	screen.HideCursor()
	screen.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorOrange))
	screen.Clear()
	for {
		screen.Fill('x', tcell.StyleDefault)
		screen.Show()
		ev := screen.PollEvent()
		if ev != nil {
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					return
				}
			}
		}
	}
}
