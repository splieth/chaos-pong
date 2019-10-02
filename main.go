package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	_ = termbox.Init()
	termbox.SetInputMode(termbox.InputEsc)
	_ = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	_ = termbox.Flush()
}
