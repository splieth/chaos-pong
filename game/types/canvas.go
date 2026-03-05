package types

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Canvas struct {
	Pos    Vector
	Center Vector
	Width  float64
	Height float64
	Color  color.Color
	Image  *ebiten.Image
	Screen *ebiten.Image
}

func NewCanvas(originPoint Vector, canvasWidth, canvasHeight float64) Canvas {
	canvasImage := ebiten.NewImage(int(canvasWidth), int(canvasHeight))
	return Canvas{
		Pos:    originPoint,
		Width:  canvasWidth,
		Height: canvasHeight,
		Center: Vector{X: canvasWidth / 2, Y: canvasHeight / 2},
		Color:  color.White,
		Image:  canvasImage,
	}
}

func (c *Canvas) Fill() {
	c.Image.Fill(c.Color)
}

func (c *Canvas) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(c.Pos.X, c.Pos.Y)
	screen.DrawImage(c.Image, &options)
}
