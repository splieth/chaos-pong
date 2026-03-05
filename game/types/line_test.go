package types

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	// given
	line1 := Line{Start: Vector{X: 0, Y: 0}, End: Vector{X: 10, Y: 10}}
	line2 := Line{Start: Vector{X: 0, Y: 10}, End: Vector{X: 10, Y: 0}}

	// when
	v, err := Intersect(line1, line2)

	// then
	assert.NoError(t, err)
	assert.Equal(t, Vector{X: 5, Y: 5}, v)
}

func TestIntersect_InverseDirection(t *testing.T) {
	// given
	line1 := Line{Start: Vector{X: 20, Y: 20}, End: Vector{X: 0, Y: 0}}
	line2 := Line{Start: Vector{X: 0, Y: 20}, End: Vector{X: 20, Y: 0}}

	// when
	v, err := Intersect(line1, line2)

	// then
	assert.NoError(t, err)
	assert.Equal(t, Vector{X: 10, Y: 10}, v)
}

func TestIntersect_Parallel(t *testing.T) {
	// given
	line1 := Line{Start: Vector{X: 0, Y: 0}, End: Vector{X: 0, Y: 1}}
	line2 := Line{Start: Vector{X: 1, Y: 0}, End: Vector{X: 1, Y: 1}}

	// when
	v, err := Intersect(line1, line2)

	// then
	assert.ErrorIs(t, err, ErrNoIntersection)
	assert.Equal(t, Vector{}, v)
}

func TestIntersect_Perpendicular(t *testing.T) {
	// given
	line1 := Line{Start: Vector{X: 0, Y: 5}, End: Vector{X: 10, Y: 5}}
	line2 := Line{Start: Vector{X: 5, Y: 0}, End: Vector{X: 5, Y: 10}}

	// when
	v, err := Intersect(line1, line2)

	// then
	assert.NoError(t, err)
	assert.Equal(t, Vector{X: 5, Y: 5}, v)
}

func TestIntersect_AtOrigin(t *testing.T) {
	// given
	line1 := Line{Start: Vector{X: -1, Y: -1}, End: Vector{X: 1, Y: 1}}
	line2 := Line{Start: Vector{X: -1, Y: 1}, End: Vector{X: 1, Y: -1}}

	// when
	v, err := Intersect(line1, line2)

	// then
	assert.NoError(t, err)
	assert.Equal(t, Vector{X: 0, Y: 0}, v)
}

func TestIntersect_NegativeCoordinates(t *testing.T) {
	// given
	line1 := Line{Start: Vector{X: -10, Y: -10}, End: Vector{X: -5, Y: -5}}
	line2 := Line{Start: Vector{X: -10, Y: -5}, End: Vector{X: -5, Y: -10}}

	// when
	v, err := Intersect(line1, line2)

	// then
	assert.NoError(t, err)
	assert.InDelta(t, -7.5, v.X, 1e-10)
	assert.InDelta(t, -7.5, v.Y, 1e-10)
}

func TestIntersect_CollinearLines(t *testing.T) {
	// given
	line1 := Line{Start: Vector{X: 0, Y: 0}, End: Vector{X: 10, Y: 10}}
	line2 := Line{Start: Vector{X: 5, Y: 5}, End: Vector{X: 15, Y: 15}}

	// when
	_, err := Intersect(line1, line2)

	// then
	assert.ErrorIs(t, err, ErrNoIntersection)
}

func TestLength(t *testing.T) {
	// given
	line := Line{Start: Vector{X: 0, Y: 0}, End: Vector{X: 3, Y: 4}}

	// when
	length := line.Length()

	// then
	assert.Equal(t, 5.0, length)
}

func TestLength_Zero(t *testing.T) {
	// given
	line := Line{Start: Vector{X: 5, Y: 5}, End: Vector{X: 5, Y: 5}}

	// when
	length := line.Length()

	// then
	assert.Equal(t, 0.0, length)
}

func TestLength_Horizontal(t *testing.T) {
	// given
	line := Line{Start: Vector{X: 0, Y: 0}, End: Vector{X: 10, Y: 0}}

	// when
	length := line.Length()

	// then
	assert.Equal(t, 10.0, length)
}

func TestDirection(t *testing.T) {
	// given
	line := Line{Start: Vector{X: 0, Y: 0}, End: Vector{X: 10, Y: 0}}

	// when
	dir := line.Direction()

	// then
	assert.InDelta(t, 1.0, dir.X, 1e-10)
	assert.InDelta(t, 0.0, dir.Y, 1e-10)
}

func TestDirection_Diagonal(t *testing.T) {
	// given
	line := Line{Start: Vector{X: 0, Y: 0}, End: Vector{X: 1, Y: 1}}

	// when
	dir := line.Direction()

	// then
	expected := 1.0 / math.Sqrt(2)
	assert.InDelta(t, expected, dir.X, 1e-10)
	assert.InDelta(t, expected, dir.Y, 1e-10)
}

func TestDirection_ZeroLength(t *testing.T) {
	// given
	line := Line{Start: Vector{X: 5, Y: 5}, End: Vector{X: 5, Y: 5}}

	// when
	dir := line.Direction()

	// then
	assert.Equal(t, Vector{X: 0, Y: 0}, dir)
}
