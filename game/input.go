package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/splieth/chaos-pong/game/types"
	"os"
)

func getPaddleMoves() (types.Vector, types.Vector) {
	leftPaddleOffset := types.Vector{X: 0, Y: 0}
	rightPaddleOffset := types.Vector{X: 0, Y: 0}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		leftPaddleOffset = types.Vector{X: 0, Y: -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		leftPaddleOffset = types.Vector{X: 0, Y: 1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		rightPaddleOffset = types.Vector{X: 0, Y: -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		rightPaddleOffset = types.Vector{X: 0, Y: 1}
	}
	return leftPaddleOffset, rightPaddleOffset
}

func handleExit() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}
