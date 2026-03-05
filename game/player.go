package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/splieth/chaos-pong/game/types"
)

type Player struct {
	side   string
	paddle *Paddle
}

type Paddle struct {
	Width, Height float64
	Color         color.Color
	types.Object
}

func NewPlayer(side string, width, height int, pos types.Vector, color color.Color, canvas *types.Canvas) Player {
	paddle := NewPaddle(width, height, pos, color, canvas)
	return Player{
		side:   side,
		paddle: &paddle,
	}
}

func NewPaddle(width, height int, pos types.Vector, color color.Color, canvas *types.Canvas) Paddle {
	image := ebiten.NewImage(width, height)
	image.Fill(color)
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

func (p *Player) Draw() {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.paddle.Pos.X, p.paddle.Pos.Y)
	p.paddle.Canvas.Image.DrawImage(p.paddle.Image, &options)
}

func (p *Player) Move(offset types.Vector) {
	if offset.Norm() > 0 {
		offset.Normalize()
		offset.Multiply(p.paddle.Velocity)
		p.paddle.Pos.Add(offset)
		p.paddle.Pos.Y = math.Max(p.paddle.Pos.Y, 0)
		p.paddle.Pos.Y = math.Min(p.paddle.Pos.Y, p.paddle.Canvas.Height-p.paddle.Height)
	}
}
