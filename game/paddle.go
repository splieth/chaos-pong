package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/splieth/chaos-pong/game/types"
	"image/color"
)

type Paddle struct {
	Width, Height float64
	Color         color.Color
	types.Object
}

func NewPaddle(width, height int, pos types.Vector, color color.Color, canvas *types.Canvas) Paddle {
	image, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	_ = image.Fill(color)
	return Paddle{
		Object: types.Object{
			Pos:    pos,
			Dir:    types.Vector{X: 0, Y: 0},
			Image:  image,
			Canvas: canvas,
		},
	}
}

func (p *Paddle) Draw() {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.Pos.X, p.Pos.Y)
	_ = p.Canvas.Image.DrawImage(p.Image, &options)
}

func (p *Paddle) Move(offset types.Vector) {
	p.Pos = p.Pos.Add(offset)
}
