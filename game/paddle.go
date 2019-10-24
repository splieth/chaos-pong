package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/splieth/chaos-pong/game/types"
	"image/color"
	"math"
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
		Width:  paddleWidth,
		Height: paddleHeight,
		Object: types.Object{
			Pos:      pos,
			Dir:      types.Vector{X: 0, Y: 0},
			Image:    image,
			Canvas:   canvas,
			Velocity: 10,
		},
	}
}

func (p *Paddle) Draw() {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.Pos.X, p.Pos.Y)
	_ = p.Canvas.Image.DrawImage(p.Image, &options)
}

func (p *Paddle) Move(offset types.Vector) {
	if offset.Norm() > 0 {
		offset.Normalize()
		offset.Multiply(p.Velocity)
		p.Pos.Add(offset)
		p.Pos.Y = math.Max(p.Pos.Y, 0)
		p.Pos.Y = math.Min(p.Pos.Y, p.Canvas.Height-p.Height)
	}
}
