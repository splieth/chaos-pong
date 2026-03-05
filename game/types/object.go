package types

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	Pos, Dir Vector
	Velocity float64
	Image    *ebiten.Image
	Canvas   *Canvas
}
