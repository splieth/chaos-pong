package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/splieth/chaos-pong/game/types"
)

func getPaddleMoves() (types.Vector, types.Vector) {
	leftPaddleOffset := types.Vector{0, 0}
	rightPaddleOffset := types.Vector{0, 0}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		leftPaddleOffset = types.Vector{0, -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		leftPaddleOffset = types.Vector{0, 1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		rightPaddleOffset = types.Vector{0, -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		rightPaddleOffset = types.Vector{0, 1}
	}
	return leftPaddleOffset, rightPaddleOffset
}
