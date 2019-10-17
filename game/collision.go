package game

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (g *Game) detectCollisions() []Collision {
	canvas := g.ballCanvas
	newPos := g.ball.GetNextPos().convertToInt()
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

func (g *Game) HandlePaddleColission(coll Collision) {
	var paddle Paddle
	var randomAddition float64
	if coll == LeftPaddle {
		paddle = *g.leftPaddle
	} else {
		paddle = *g.rightPaddle
	}
	lastPaddleDir := paddle.lastDirection.y
	randomAddition = randomize(paddle.lastDirection)
	g.ball.direction.x = g.ball.direction.x * -1
	log.Println("Direction before random addition: %f", g.ball.direction.y)
	log.Println("Last direction before random addition: %f", g.ball.direction.y)
	g.ball.direction.y = float64(lastPaddleDir) + randomAddition
	log.Println("New direction with random addition: %f", g.ball.direction.y)
}

func randomize(v IntVector) float64 {
	var randomAddition float64
	if v.y > 0 {
		randomAddition = -float64(rand.Intn(1000)) / 1000
	} else {
		randomAddition = float64(rand.Intn(1000)) / 1000
	}
	log.Printf("Got randomAddition: %f", randomAddition)
	return randomAddition
}

func (g *Game) HandleBallCollision() {
	collisions := g.detectCollisions()
	for _, coll := range collisions {
		switch coll {
		case TopWall, BottomWall:
			g.ball.direction.y = g.ball.direction.y * -1
		case RightPaddle, LeftPaddle:
			g.HandlePaddleColission(coll)
			g.increaseBallSpeed()
		case RightWall:
			g.scoreGoal(RightWall)
		case LeftWall:
			g.scoreGoal(LeftWall)
		}
	}
}
