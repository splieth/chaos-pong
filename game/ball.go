package game

import (
	"github.com/gdamore/tcell"
)

type Ball struct {
	position  Vector
	direction Vector
	color     tcell.Color
}

type Collision int8

const (
	TopWall Collision = iota
	RightWall
	BottomWall
	LeftWall
	LeftPaddle
	RightPaddle
)

func (ball *Ball) Move() {
	ball.position = Add(ball.position, ball.direction)
}

func (ball *Ball) Draw(screen tcell.Screen) {
	screen.SetContent(ball.position.x, ball.position.y, '●', nil,
		tcell.StyleDefault.Background(ballCanvasBG).Foreground(ball.color))
}

func (ball *Ball) detectCollisions(g *Game) []Collision {
	canvas := g.ballCanvas
	newPos := Add(ball.position, ball.direction)
	var collisions []Collision
	if newPos.x == g.leftPaddle.position.x &&
		newPos.y >= g.leftPaddle.position.y &&
		newPos.y <= g.leftPaddle.position.y+g.leftPaddle.height {
		collisions = append(collisions, LeftPaddle)
	}
	if newPos.x == g.rightPaddle.position.x &&
		newPos.y >= g.rightPaddle.position.y &&
		newPos.y <= g.rightPaddle.position.y+g.rightPaddle.height {
		collisions = append(collisions, RightPaddle)
	}
	if newPos.x < canvas.x {
		collisions = append(collisions, LeftWall)
	}
	if newPos.x >= canvas.x+canvas.width {
		collisions = append(collisions, RightWall)
	}
	if newPos.y < canvas.y {
		collisions = append(collisions, TopWall)
	}
	if newPos.y >= canvas.y+canvas.height {
		collisions = append(collisions, BottomWall)
	}
	return collisions
}

func (ball *Ball) HandleCollision(g *Game) {
	collisions := ball.detectCollisions(g)
	for _, coll := range collisions {
		switch coll {
		case TopWall, BottomWall:
			ball.direction.y = ball.direction.y * -1
		case RightPaddle:
			ball.direction.x = ball.direction.x * -1
			ball.direction.y = g.rightPaddle.lastDirection.y
		case LeftPaddle:
			ball.direction.x = ball.direction.x * -1
			ball.direction.y = g.leftPaddle.lastDirection.y
		case RightWall:
			g.scores[0] += 1
			g.goooooooal()
		case LeftWall:
			g.scores[1] += 1
			g.goooooooal()
		}
	}
}

func (b *Ball) center(g *Game) {
	b.position.x = (g.ballCanvas.width / 2) + canvasPadding
	b.position.y = (g.ballCanvas.height / 2) + canvasPadding
}
