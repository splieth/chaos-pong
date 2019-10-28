package types

import (
	"fmt"
	"math"
	"math/rand"
)

type Vector struct {
	X, Y float64
}

type Line struct {
	Start, End Vector
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

func (l *Line) diff() (float64, float64) {
	return l.Start.X - l.End.X, l.Start.Y - l.End.Y
}

func Intersect(line1, line2 Line) (Vector, error) {
	xDiff1, yDiff1 := line1.diff()
	xDiff2, yDiff2 := line2.diff()
	xDiff := Vector{xDiff1, xDiff2}
	yDiff := Vector{yDiff1, yDiff2}
	div := det(xDiff, yDiff)
	if div == 0 {
		return Vector{}, fmt.Errorf("No intersection")
	}
	d := Vector{det(line1.Start, line1.End), det(line2.Start, line2.End)}
	return Vector{det(d, xDiff) / div, det(d, yDiff) / div}, nil
}
