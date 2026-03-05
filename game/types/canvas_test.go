package types

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	// given
	pos := Vector{X: 50, Y: 100}
	width := 800.0
	height := 600.0

	// when
	canvas := NewCanvas(pos, width, height)

	// then
	assert.Equal(t, pos, canvas.Pos)
	assert.Equal(t, width, canvas.Width)
	assert.Equal(t, height, canvas.Height)
	assert.Equal(t, Vector{X: 400, Y: 300}, canvas.Center)
	assert.Equal(t, color.White, canvas.Color)
	assert.NotNil(t, canvas.Image)
}

func TestNewCanvas_ImageDimensions(t *testing.T) {
	// given
	pos := Vector{X: 0, Y: 0}

	// when
	canvas := NewCanvas(pos, 320, 240)

	// then
	w, h := canvas.Image.Size()
	assert.Equal(t, 320, w)
	assert.Equal(t, 240, h)
}

func TestNewCanvas_Center(t *testing.T) {
	// given
	pos := Vector{X: 10, Y: 20}

	// when
	canvas := NewCanvas(pos, 100, 50)

	// then
	assert.Equal(t, Vector{X: 50, Y: 25}, canvas.Center)
}

func TestCanvas_Fill(t *testing.T) {
	// given
	canvas := NewCanvas(Vector{X: 0, Y: 0}, 10, 10)
	canvas.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	// when / then — should not panic
	assert.NotPanics(t, func() {
		canvas.Fill()
	})
}

func TestCanvas_Draw(t *testing.T) {
	// given
	canvas := NewCanvas(Vector{X: 10, Y: 20}, 50, 50)
	canvas.Fill()
	screen := NewCanvas(Vector{X: 0, Y: 0}, 100, 100)

	// when / then — should not panic
	assert.NotPanics(t, func() {
		canvas.Draw(screen.Image)
	})
}

func TestCanvas_Contains_Inside(t *testing.T) {
	// given
	canvas := NewCanvas(Vector{X: 0, Y: 0}, 100, 50)
	point := Vector{X: 50, Y: 25}

	// when
	result := canvas.Contains(point)

	// then
	assert.True(t, result)
}

func TestCanvas_Contains_Origin(t *testing.T) {
	// given
	canvas := NewCanvas(Vector{X: 0, Y: 0}, 100, 50)
	point := Vector{X: 0, Y: 0}

	// when
	result := canvas.Contains(point)

	// then
	assert.True(t, result)
}

func TestCanvas_Contains_Edge(t *testing.T) {
	// given
	canvas := NewCanvas(Vector{X: 0, Y: 0}, 100, 50)
	point := Vector{X: 100, Y: 50}

	// when
	result := canvas.Contains(point)

	// then
	assert.True(t, result)
}

func TestCanvas_Contains_OutsideLeft(t *testing.T) {
	// given
	canvas := NewCanvas(Vector{X: 0, Y: 0}, 100, 50)

	// when
	result := canvas.Contains(Vector{X: -1, Y: 25})

	// then
	assert.False(t, result)
}

func TestCanvas_Contains_OutsideRight(t *testing.T) {
	// given
	canvas := NewCanvas(Vector{X: 0, Y: 0}, 100, 50)

	// when
	result := canvas.Contains(Vector{X: 101, Y: 25})

	// then
	assert.False(t, result)
}

func TestCanvas_Contains_OutsideTop(t *testing.T) {
	// given
	canvas := NewCanvas(Vector{X: 0, Y: 0}, 100, 50)

	// when
	result := canvas.Contains(Vector{X: 50, Y: -1})

	// then
	assert.False(t, result)
}

func TestCanvas_Contains_OutsideBottom(t *testing.T) {
	// given
	canvas := NewCanvas(Vector{X: 0, Y: 0}, 100, 50)

	// when
	result := canvas.Contains(Vector{X: 50, Y: 51})

	// then
	assert.False(t, result)
}
