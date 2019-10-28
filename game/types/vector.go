package types

import (
	"math"
	"math/rand"
)

type Vector struct {
	X, Y float64
}

func (v *Vector) Add(other Vector) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vector) Multiply(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
}

func (v *Vector) InvertX() {
	v.X *= -1
}

func (v *Vector) InvertY() {
	v.Y *= -1
}

func (v *Vector) Norm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() {
	norm := v.Norm()
	v.X = v.X / norm
	v.Y = v.Y / norm
}

func (v *Vector) Randomize() {
	var randomAddition float64
	if v.Y > 0 {
		randomAddition = -float64(rand.Intn(1000)) / 1000
	} else {
		randomAddition = float64(rand.Intn(1000)) / 1000
	}
	v.Y += randomAddition
}

func det(a, b Vector) float64 {
	return a.X*b.Y - a.Y*b.X
}
