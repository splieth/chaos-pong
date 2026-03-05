package game

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/splieth/chaos-pong/game/types"
	"github.com/stretchr/testify/assert"
)

// newTestGame creates a minimal game for testing without loading assets from disk.
func newTestGame() *Game {
	ballCanvas := types.NewCanvas(types.Vector{X: 0, Y: 0}, 800, 600)
	scoreCanvas := types.NewCanvas(types.Vector{X: 0, Y: 600}, 800, 50)
	scoreCanvas.Color = color.Black

	ballImage := ebiten.NewImage(10, 10)
	ball := Ball{
		Radius:   5,
		Diameter: 10,
		Object: types.Object{
			Pos:      ballCanvas.Center,
			Dir:      types.Vector{X: 1, Y: 1},
			Image:    ballImage,
			Canvas:   &ballCanvas,
			Velocity: InitialBallSpeed,
		},
	}

	leftPos := types.Vector{X: 0, Y: 0}
	rightPos := types.Vector{X: ballCanvas.Width - paddleWidth, Y: 0}

	player := NewPlayer("left", paddleWidth, paddleHeight, leftPos, color.RGBA{R: 255, A: 255}, &ballCanvas)
	npc := NewPlayer("right", paddleWidth, paddleHeight, rightPos, color.RGBA{G: 255, A: 255}, &ballCanvas)

	return &Game{
		screenWidth:  800,
		screenHeight: 650,
		ball:         &ball,
		ballCanvas:   &ballCanvas,
		scoreCanvas:  &scoreCanvas,
		player:       &player,
		npc:          &npc,
		score: map[string]int{
			player1: 0,
			player2: 0,
		},
		started: true,
	}
}

// --- Wall collision tests ---

func TestHandleBallCanvasCollision_TopWall(t *testing.T) {
	// given
	g := newTestGame()
	g.ball.Pos = types.Vector{X: 400, Y: -1}
	g.ball.Dir = types.Vector{X: 1, Y: -1}

	// when
	wall := g.handleBallCanvasCollision()

	// then
	assert.Equal(t, TopWall, wall)
	assert.True(t, g.ball.Dir.Y > 0)
}

func TestHandleBallCanvasCollision_BottomWall(t *testing.T) {
	// given
	g := newTestGame()
	g.ball.Pos = types.Vector{X: 400, Y: 600}
	g.ball.Dir = types.Vector{X: 1, Y: 1}

	// when
	wall := g.handleBallCanvasCollision()

	// then
	assert.Equal(t, BottomWall, wall)
	assert.True(t, g.ball.Dir.Y < 0)
}

func TestHandleBallCanvasCollision_RightWall(t *testing.T) {
	// given
	g := newTestGame()
	g.ball.Pos = types.Vector{X: 800, Y: 300}
	g.ball.Dir = types.Vector{X: 1, Y: 1}

	// when
	wall := g.handleBallCanvasCollision()

	// then
	assert.Equal(t, RightWall, wall)
	assert.True(t, g.ball.Dir.X < 0)
}

func TestHandleBallCanvasCollision_LeftWall(t *testing.T) {
	// given
	g := newTestGame()
	g.ball.Pos = types.Vector{X: -1, Y: 300}
	g.ball.Dir = types.Vector{X: -1, Y: 1}

	// when
	wall := g.handleBallCanvasCollision()

	// then
	assert.Equal(t, LeftWall, wall)
	assert.True(t, g.ball.Dir.X > 0)
}

func TestHandleBallCanvasCollision_NoWall(t *testing.T) {
	// given
	g := newTestGame()
	g.ball.Pos = types.Vector{X: 400, Y: 300}
	g.ball.Dir = types.Vector{X: 1, Y: 1}

	// when
	wall := g.handleBallCanvasCollision()

	// then
	assert.Equal(t, NoWall, wall)
}

// --- Score tests ---

func TestHandleScores_RightWallScoresPlayer1(t *testing.T) {
	// given
	g := newTestGame()

	// when
	g.handleScores(RightWall)

	// then
	assert.Equal(t, 1, g.score[player1])
	assert.Equal(t, 0, g.score[player2])
}

func TestHandleScores_LeftWallScoresPlayer2(t *testing.T) {
	// given
	g := newTestGame()

	// when
	g.handleScores(LeftWall)

	// then
	assert.Equal(t, 0, g.score[player1])
	assert.Equal(t, 1, g.score[player2])
}

func TestHandleScores_TopWallNoScore(t *testing.T) {
	// given
	g := newTestGame()

	// when
	g.handleScores(TopWall)

	// then
	assert.Equal(t, 0, g.score[player1])
	assert.Equal(t, 0, g.score[player2])
}

func TestHandleScores_NoWallNoScore(t *testing.T) {
	// given
	g := newTestGame()

	// when
	g.handleScores(NoWall)

	// then
	assert.Equal(t, 0, g.score[player1])
	assert.Equal(t, 0, g.score[player2])
}

// --- Paddle collision tests ---

func TestHandleBallPaddleCollision_HitsLeftPaddle(t *testing.T) {
	// given
	g := newTestGame()
	g.player.paddle.Pos = types.Vector{X: 0, Y: 100}
	g.ball.Pos = types.Vector{X: 20, Y: 150}
	g.ball.Dir = types.Vector{X: -1, Y: 0}
	originalVelocity := g.ball.Velocity

	// when
	g.handleBallPaddleCollision()

	// then
	assert.True(t, g.ball.Dir.X > 0)
	assert.Equal(t, originalVelocity+1, g.ball.Velocity)
}

func TestHandleBallPaddleCollision_HitsRightPaddle(t *testing.T) {
	// given
	g := newTestGame()
	g.npc.paddle.Pos = types.Vector{X: 770, Y: 100}
	g.ball.Pos = types.Vector{X: 765, Y: 150}
	g.ball.Dir = types.Vector{X: 1, Y: 0}
	originalVelocity := g.ball.Velocity

	// when
	g.handleBallPaddleCollision()

	// then
	assert.True(t, g.ball.Dir.X < 0)
	assert.Equal(t, originalVelocity+1, g.ball.Velocity)
}

func TestHandleBallPaddleCollision_MissesPaddle(t *testing.T) {
	// given
	g := newTestGame()
	g.player.paddle.Pos = types.Vector{X: 0, Y: 100}
	g.ball.Pos = types.Vector{X: 400, Y: 300}
	g.ball.Dir = types.Vector{X: 1, Y: 0}
	originalVelocity := g.ball.Velocity

	// when
	g.handleBallPaddleCollision()

	// then
	assert.InDelta(t, 1.0, g.ball.Dir.X, 1e-10)
	assert.Equal(t, originalVelocity, g.ball.Velocity)
}

// --- Layout test ---

func TestLayout(t *testing.T) {
	// given
	g := newTestGame()

	// when
	w, h := g.Layout(1920, 1080)

	// then
	assert.Equal(t, 800, w)
	assert.Equal(t, 650, h)
}

// --- StartGame test ---

func TestStartGame(t *testing.T) {
	// given
	g := newTestGame()
	g.started = false

	// when
	g.StartGame("left")

	// then
	assert.True(t, g.started)
	assert.Equal(t, "left", g.player.side)
}
