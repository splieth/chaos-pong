package types

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Paddle struct {
	Width, Height float64
	Color         color.Color
	Object
}

func NewPaddle(startPosition, direction Vector, canvas *Canvas, image *ebiten.Image) Paddle {
	return Paddle{
		Object: Object{
			pos:    startPosition,
			dir:    direction,
			image:  image,
			canvas: canvas,
		},
	}
}

func (p *Paddle) Draw() {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.pos.X, p.pos.Y)
	_ = p.canvas.Image.DrawImage(p.image, &options)
}

func (p *Paddle) Move(offset Vector) {
	p.pos = p.pos.Add(offset)
}
