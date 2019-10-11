package game

import "github.com/gdamore/tcell"

type Ball struct {
	sprite Sprite
}

func (ball *Ball) getNextPos() Vector {
	newX := ball.sprite.position.x + ball.sprite.direction.x
	newY := ball.sprite.position.y + ball.sprite.direction.y
	return Vector{newX, newY}
}

func (ball *Ball) move(screen tcell.Screen) {
	ball.sprite.position = ball.getNextPos()
	screen.SetContent(ball.sprite.position.x, ball.sprite.position.y, '●', nil, tcell.StyleDefault.Background(tcell.ColorRebeccaPurple).Foreground(tcell.ColorOrangeRed))
}
