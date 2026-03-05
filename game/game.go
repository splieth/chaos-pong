package game

import (
	"image/color"
	"log"
	"strconv"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/splieth/chaos-pong/game/types"
	"golang.org/x/image/font"
)

const (
	canvasPadding     = 50.0
	scoreCanvasHeight = 50
	paddleWidth       = 25
	paddleHeight      = 150
	player1           = "player1"
	player2           = "player2"
)

// Game holds the complete state of a Chaos Pong match, including the ball,
// players, canvases, and scores. It implements the ebiten.Game interface.
type Game struct {
	screenWidth  int
	screenHeight int
	ball         *Ball
	ballCanvas   *types.Canvas
	scoreCanvas  *types.Canvas
	score        map[string]int
	scoreFont    font.Face
	started      bool
	player       *Player
	npc          *Player
}

// NewGame creates a fully initialized game with the given screen dimensions.
// The ball is placed at the center of the play area and both players are
// positioned at their respective sides.
func NewGame(screenWidth, screenHeight int, basePath string) Game {
	ballCanvas := createBallCanvas(screenWidth, screenHeight)
	scoreCanvas := createScoreCanvas(screenWidth, ballCanvas.Height+10+canvasPadding)

	ball := newBall(&ballCanvas, basePath)

	tt, _ := truetype.Parse(fonts.MPlus1pRegular_ttf)
	scoreFont := truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	leftPaddlePos := types.Vector{X: 0, Y: 0}
	rightPaddlePos := types.Vector{X: ballCanvas.Width - paddleWidth, Y: 0}

	leftColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	rightColor := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	player := NewPlayer("left", paddleWidth, paddleHeight, leftPaddlePos, leftColor, &ballCanvas)
	npc := NewPlayer("right", paddleWidth, paddleHeight, rightPaddlePos, rightColor, &ballCanvas)

	return Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		ball:         &ball,
		ballCanvas:   &ballCanvas,
		scoreCanvas:  &scoreCanvas,
		player:       &player,
		npc:          &npc,
		score: map[string]int{
			player1: 0,
			player2: 0,
		},
		scoreFont: scoreFont,
		started:   true,
	}
}

// createBallCanvas builds the main play area canvas, inset from the screen
// edges by canvasPadding on all sides.
func createBallCanvas(screenWidth, screenHeight int) types.Canvas {
	width := float64(screenWidth) - 2*canvasPadding
	height := float64(screenHeight) - 2*canvasPadding - scoreCanvasHeight
	return types.NewCanvas(types.Vector{X: canvasPadding, Y: canvasPadding}, width, height)
}

// createScoreCanvas builds the score display canvas below the play area.
func createScoreCanvas(screenWidth int, yCoordinate float64) types.Canvas {
	width := float64(screenWidth) - 2*canvasPadding
	canvas := types.NewCanvas(types.Vector{X: canvasPadding, Y: yCoordinate}, width, scoreCanvasHeight)
	canvas.Color = color.Black
	return canvas
}

// Update advances the game state by one tick. It processes input, handles
// collisions, updates scores, and moves the ball.
func (g *Game) Update() error {
	handleExit()
	if !g.started {
		return nil
	}
	collidedWall := g.handleBallCanvasCollision()
	leftPaddleOffset, rightPaddleOffset := getPaddleMoves()
	g.player.Move(leftPaddleOffset)
	g.npc.Move(rightPaddleOffset)
	g.handleScores(collidedWall)
	g.handleBallPaddleCollision()
	g.ball.Move()
	return nil
}

// Draw renders the current game state onto the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.ballCanvas.Fill()
	g.scoreCanvas.Fill()
	g.drawScores()
	g.ball.Draw()
	g.player.Draw()
	g.npc.Draw()
	g.ballCanvas.Draw(screen)
	g.scoreCanvas.Draw(screen)
}

// Layout returns the logical screen dimensions for ebiten.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

// handleScores increments the appropriate player's score when the ball
// hits a side wall and resets the ball to center.
func (g *Game) handleScores(wall Wall) {
	switch wall {
	case RightWall:
		g.score[player1]++
		g.reset()
	case LeftWall:
		g.score[player2]++
		g.reset()
	}
}

// drawScores renders the current score text onto the score canvas.
func (g *Game) drawScores() {
	score := strconv.Itoa(g.score[player1]) + ":" + strconv.Itoa(g.score[player2])
	x := int(g.scoreCanvas.Width)/2 - 2*len(score)
	text.Draw(g.scoreCanvas.Image, score, g.scoreFont, x, 25, color.White)
}

// reset pauses briefly, then moves the ball back to center and restores
// its initial speed.
func (g *Game) reset() {
	time.Sleep(1 * time.Second)
	g.ball.Pos = g.ballCanvas.Center
	g.ball.Velocity = InitialBallSpeed
}

// LoadImage loads an image file from disk and returns an ebiten image.
// Panics on failure since missing assets are unrecoverable.
func LoadImage(resourcesBasePath, path string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(resourcesBasePath + path)
	if err != nil {
		log.Fatal(err)
	}
	return image
}

// StartGame activates the game for the given side. Used by the multiplayer
// client after receiving the start signal from the server.
func (g *Game) StartGame(side string) {
	p := Player{side: side}
	g.player = &p
	g.started = true
}
