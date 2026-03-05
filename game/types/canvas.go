package types

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Canvas represents a drawable rectangular area with its own off-screen image.
// Game objects draw onto the canvas image, which is then composited onto
// the main screen at the canvas position.
type Canvas struct {
	Pos    Vector
	Center Vector
	Width  float64
	Height float64
	Color  color.Color
	Image  *ebiten.Image
}

// NewCanvas creates a canvas at the given position with the specified dimensions.
// The canvas is initialized with a white background color.
func NewCanvas(pos Vector, width, height float64) Canvas {
	return Canvas{
		Pos:    pos,
		Width:  width,
		Height: height,
		Center: Vector{X: width / 2, Y: height / 2},
		Color:  color.White,
		Image:  ebiten.NewImage(int(width), int(height)),
	}
}

// Fill clears the canvas image and fills it with the canvas color.
func (c *Canvas) Fill() {
	c.Image.Fill(c.Color)
}

// Draw composites the canvas image onto the screen at the canvas position.
func (c *Canvas) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(c.Pos.X, c.Pos.Y)
	screen.DrawImage(c.Image, opts)
}

// Contains reports whether the given point lies within the canvas bounds.
// Coordinates are relative to the canvas (not the screen).
func (c *Canvas) Contains(point Vector) bool {
	return point.X >= 0 && point.X <= c.Width &&
		point.Y >= 0 && point.Y <= c.Height
}
