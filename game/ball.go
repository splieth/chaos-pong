package game

import "github.com/gdamore/tcell"

type Ball struct {
	sprite Sprite
	speed  int
}

func (ball *Ball) getNextPos() Vector {
	newX := ball.sprite.position.x + ball.sprite.direction.x*ball.speed
	newY := ball.sprite.position.y + ball.sprite.direction.y*ball.speed
	return Vector{newX, newY}
}

func (ball *Ball) move(screen tcell.Screen) {
	ball.sprite.position = ball.getNextPos()
	screen.SetContent(ball.sprite.position.x, ball.sprite.position.y, '●', nil, tcell.StyleDefault.Background(tcell.ColorRebeccaPurple).Foreground(tcell.ColorOrangeRed))
}
