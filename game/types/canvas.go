package types

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Canvas struct {
	X, Y  float64
	Color color.Color
	Image *ebiten.Image
}

func (c *Canvas) Fill() {
	_ = c.Image.Fill(c.Color)
}
