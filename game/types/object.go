package types

import (
	"github.com/hajimehoshi/ebiten"
)

type Object struct {
	Pos, Dir Vector
	Image    *ebiten.Image
	Canvas   *Canvas
}
