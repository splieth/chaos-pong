package game

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/splieth/chaos-pong/game/types"
	"golang.org/x/image/font"
	"image/color"
	"log"
	"strconv"
	"time"
)

const (
	canvasPadding     = 50.0
	scoreCanvasHeight = 50
	paddleWidth       = 25
	paddleHeight      = 150
	player1           = "player1"
	player2           = "player2"
)

type Game struct {
	screenWidth  int
	screenHeight int
	ball        *Ball
	ballCanvas  *types.Canvas
	scoreCanvas *types.Canvas
	score       map[string]int
	scoreFont    font.Face
	started      bool
	player       *Player
	npc          *Player
}

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

	leftPaddleColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	rightPaddleColor := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	leftPaddlePos := types.Vector{X: 0, Y: 0}
	rightPaddlePos := types.Vector{X: ballCanvas.Width - paddleWidth, Y: 0}

	player := NewPlayer("left", paddleWidth, paddleHeight, leftPaddlePos, leftPaddleColor, &ballCanvas)
	npc := NewPlayer("right", paddleWidth, paddleHeight, rightPaddlePos, rightPaddleColor, &ballCanvas)

	return Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		ball:        &ball,
		ballCanvas:  &ballCanvas,
		scoreCanvas: &scoreCanvas,
		player:      &player,
		npc:          &npc,
		score: map[string]int{
			player1: 0,
			player2: 0,
		},
		scoreFont: scoreFont,
		started:   true,
	}
}

func createBallCanvas(screenWidth, screenHeight int) types.Canvas {
	canvasWidth := float64(screenWidth) - 2*canvasPadding
	canvasHeight := float64(screenHeight) - 2*canvasPadding - scoreCanvasHeight
	ballCanvas := types.NewCanvas(types.Vector{X: canvasPadding, Y: canvasPadding}, canvasWidth, canvasHeight)
	return ballCanvas
}

func createScoreCanvas(screenWidth int, yCoordinate float64) types.Canvas {
	canvasWidth := float64(screenWidth) - 2*canvasPadding
	scoreCanvas := types.NewCanvas(types.Vector{X: canvasPadding, Y: yCoordinate}, canvasWidth, scoreCanvasHeight)
	scoreCanvas.Color = color.Black
	return scoreCanvas
}

func (g *Game) Update() error {
	handleExit()
	if g.started {
		collidedWall := g.handleBallCanvasCollision()
		leftPaddleOffset, rightPaddleOffset := getPaddleMoves()
		g.player.Move(leftPaddleOffset)
		g.npc.Move(rightPaddleOffset)
		g.handleScores(collidedWall)
		g.handleBallPaddleCollision()
		g.ball.Move()
	}
	return nil
}

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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) handleScores(wall Wall) {
	if wall == RightWall {
		g.score[player1]++
		g.reset()
	}
	if wall == LeftWall {
		g.score[player2]++
		g.reset()
	}
}

func (g *Game) drawScores() {
	score := strconv.Itoa(g.score[player1]) + ":" + strconv.Itoa(g.score[player2])
	text.Draw(g.scoreCanvas.Image, score, g.scoreFont, int(g.scoreCanvas.Width)/2-2*len(score), 25, color.White)
}

func (g *Game) reset() {
	time.Sleep(1 * time.Second)
	g.ball.Pos = g.ballCanvas.Center
	g.ball.Velocity = InitialBallSpeed
}

func LoadImage(resourcesBasePath, path string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(resourcesBasePath + path)
	if err != nil {
		log.Fatal(err)
	}
	return image
}

func (g *Game) StartGame(side string) {
	p := Player{side: side}
	g.player = &p
	g.started = true
}
