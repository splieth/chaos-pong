package game

import "github.com/gdamore/tcell"

const (
	ballColor = tcell.ColorOrangeRed
)

type Ball struct {
	sprite Sprite
}

func (ball *Ball) getNextPos() Vector {
	newX := ball.sprite.position.x + ball.sprite.direction.x
	newY := ball.sprite.position.y + ball.sprite.direction.y
	return Vector{newX, newY}
}

func (ball *Ball) move() {
	ball.sprite.position = ball.getNextPos()
}

func (ball *Ball) draw(screen tcell.Screen) {
	screen.SetContent(ball.sprite.position.x, ball.sprite.position.y, '●', nil, tcell.StyleDefault.Background(backgroundColor).Foreground(ballColor))
}
