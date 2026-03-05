package game

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/splieth/chaos-pong/game/types"
)

var (
	up   = types.Vector{X: 0, Y: -1}
	down = types.Vector{X: 0, Y: 1}
)

// getPaddleMoves reads keyboard state and returns directional offset vectors
// for the left paddle (W/S keys) and right paddle (Up/Down keys).
func getPaddleMoves() (types.Vector, types.Vector) {
	var left, right types.Vector

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		left = up
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		left = down
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		right = up
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		right = down
	}

	return left, right
}

// handleExit terminates the process when Escape is pressed.
func handleExit() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}
