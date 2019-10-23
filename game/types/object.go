package types

import (
	"github.com/hajimehoshi/ebiten"
)

type Object struct {
	pos, dir Vector
	image    *ebiten.Image
	canvas   *Canvas
}
