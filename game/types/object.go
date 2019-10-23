package types

import (
	"github.com/hajimehoshi/ebiten"
)

type Object struct {
	pos, vel Vector
	image    *ebiten.Image
	canvas   *ebiten.Image
}
