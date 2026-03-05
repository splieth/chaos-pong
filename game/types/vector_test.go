package types

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// given
	v := Vector{X: 1, Y: 2}

	// when
	v.Add(Vector{X: 3, Y: 4})

	// then
	assert.Equal(t, Vector{X: 4, Y: 6}, v)
}

func TestAdd_Negative(t *testing.T) {
	// given
	v := Vector{X: 5, Y: 5}

	// when
	v.Add(Vector{X: -3, Y: -7})

	// then
	assert.Equal(t, Vector{X: 2, Y: -2}, v)
}

func TestMultiply(t *testing.T) {
	// given
	v := Vector{X: 3, Y: 4}

	// when
	v.Multiply(2)

	// then
	assert.Equal(t, Vector{X: 6, Y: 8}, v)
}

func TestMultiply_Zero(t *testing.T) {
	// given
	v := Vector{X: 3, Y: 4}

	// when
	v.Multiply(0)

	// then
	assert.Equal(t, Vector{X: 0, Y: 0}, v)
}

func TestMultiply_Fractional(t *testing.T) {
	// given
	v := Vector{X: 10, Y: 4}

	// when
	v.Multiply(0.5)

	// then
	assert.Equal(t, Vector{X: 5, Y: 2}, v)
}

func TestInvertX(t *testing.T) {
	// given
	v := Vector{X: 3, Y: 4}

	// when
	v.InvertX()

	// then
	assert.Equal(t, Vector{X: -3, Y: 4}, v)
}

func TestInvertX_AlreadyNegative(t *testing.T) {
	// given
	v := Vector{X: -3, Y: 4}

	// when
	v.InvertX()

	// then
	assert.Equal(t, Vector{X: 3, Y: 4}, v)
}

func TestInvertY(t *testing.T) {
	// given
	v := Vector{X: 3, Y: 4}

	// when
	v.InvertY()

	// then
	assert.Equal(t, Vector{X: 3, Y: -4}, v)
}

func TestNorm(t *testing.T) {
	// given
	v := Vector{X: 3, Y: 4}

	// when
	norm := v.Norm()

	// then
	assert.Equal(t, 5.0, norm)
}

func TestNorm_Zero(t *testing.T) {
	// given
	v := Vector{X: 0, Y: 0}

	// when
	norm := v.Norm()

	// then
	assert.Equal(t, 0.0, norm)
}

func TestNorm_UnitX(t *testing.T) {
	// given
	v := Vector{X: 1, Y: 0}

	// when
	norm := v.Norm()

	// then
	assert.Equal(t, 1.0, norm)
}

func TestNormalize(t *testing.T) {
	// given
	v := Vector{X: 3, Y: 4}

	// when
	v.Normalize()

	// then
	assert.InDelta(t, 0.6, v.X, 1e-10)
	assert.InDelta(t, 0.8, v.Y, 1e-10)
	assert.InDelta(t, 1.0, v.Norm(), 1e-10)
}

func TestNormalize_ZeroVector(t *testing.T) {
	// given
	v := Vector{X: 0, Y: 0}

	// when
	v.Normalize()

	// then
	assert.Equal(t, Vector{X: 0, Y: 0}, v)
}

func TestNormalize_AlreadyUnit(t *testing.T) {
	// given
	v := Vector{X: 1, Y: 0}

	// when
	v.Normalize()

	// then
	assert.InDelta(t, 1.0, v.X, 1e-10)
	assert.InDelta(t, 0.0, v.Y, 1e-10)
}

func TestNormalize_Diagonal(t *testing.T) {
	// given
	v := Vector{X: 1, Y: 1}

	// when
	v.Normalize()

	// then
	expected := 1.0 / math.Sqrt(2)
	assert.InDelta(t, expected, v.X, 1e-10)
	assert.InDelta(t, expected, v.Y, 1e-10)
}

func TestRandomize_PositiveY(t *testing.T) {
	// given
	v := Vector{X: 1, Y: 0.5}
	originalY := v.Y

	// when
	v.Randomize()

	// then
	assert.LessOrEqual(t, v.Y, originalY)
	assert.Equal(t, 1.0, v.X)
}

func TestRandomize_NegativeY(t *testing.T) {
	// given
	v := Vector{X: 1, Y: -0.5}
	originalY := v.Y

	// when
	v.Randomize()

	// then
	assert.GreaterOrEqual(t, v.Y, originalY)
	assert.Equal(t, 1.0, v.X)
}

func TestDet(t *testing.T) {
	// given
	a := Vector{X: 1, Y: 2}
	b := Vector{X: 3, Y: 4}

	// when
	result := det(a, b)

	// then
	assert.Equal(t, -2.0, result)
}

func TestDet_Parallel(t *testing.T) {
	// given
	a := Vector{X: 1, Y: 2}
	b := Vector{X: 2, Y: 4}

	// when
	result := det(a, b)

	// then
	assert.Equal(t, 0.0, result)
}
