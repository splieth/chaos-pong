package types

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Canvas struct {
	X, Y   float64
	Color  color.Color
	Image  *ebiten.Image
	Screen *ebiten.Image
}

func NewCanvas(screen *ebiten.Image, x, y, padding float64) Canvas {
	screenWidth, screenHeight := screen.Size()
	canvasWidth := float64(screenWidth) - 2*padding
	canvasHeight := float64(screenHeight) - 2*padding
	canvasImage, _ := ebiten.NewImage(int(canvasWidth), int(canvasHeight), ebiten.FilterDefault)
	return Canvas{
		X:     x,
		Y:     y,
		Color: color.White,
		Image: canvasImage,
	}
}

func (c *Canvas) Fill() {
	_ = c.Image.Fill(c.Color)
}

func (c *Canvas) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(c.X, c.Y)
	_ = screen.DrawImage(c.Image, &options)
}
