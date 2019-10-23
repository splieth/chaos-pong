package types

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Canvas struct {
	Pos    Vector
	Center Vector
	Color  color.Color
	Image  *ebiten.Image
	Screen *ebiten.Image
}

func NewCanvas(screen *ebiten.Image, originPoint Vector, padding float64) Canvas {
	screenWidth, screenHeight := screen.Size()
	canvasWidth := float64(screenWidth) - 2*padding
	canvasHeight := float64(screenHeight) - 2*padding
	canvasImage, _ := ebiten.NewImage(int(canvasWidth), int(canvasHeight), ebiten.FilterDefault)
	return Canvas{
		Pos:    originPoint,
		Center: Vector{X: canvasWidth / 2, Y: canvasHeight / 2},
		Color:  color.White,
		Image:  canvasImage,
	}
}

func (c *Canvas) Fill() {
	_ = c.Image.Fill(c.Color)
}

func (c *Canvas) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(c.Pos.X, c.Pos.Y)
	_ = screen.DrawImage(c.Image, &options)
}
