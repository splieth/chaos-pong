package types

import (
	"math"
	"math/rand"
)

// Vector represents a 2D vector with X and Y components.
type Vector struct {
	X, Y float64
}

// Add adds another vector to this vector in place.
func (v *Vector) Add(other Vector) {
	v.X += other.X
	v.Y += other.Y
}

// Multiply scales the vector by a scalar value in place.
func (v *Vector) Multiply(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
}

// InvertX negates the X component, reflecting the vector horizontally.
func (v *Vector) InvertX() {
	v.X = -v.X
}

// InvertY negates the Y component, reflecting the vector vertically.
func (v *Vector) InvertY() {
	v.Y = -v.Y
}

// Norm returns the magnitude (length) of the vector.
func (v *Vector) Norm() float64 {
	return math.Hypot(v.X, v.Y)
}

// Normalize scales the vector to unit length. If the vector is zero-length,
// it is left unchanged to avoid division by zero.
func (v *Vector) Normalize() {
	norm := v.Norm()
	if norm == 0 {
		return
	}
	v.X /= norm
	v.Y /= norm
}

// Randomize adds a random perturbation to the Y component, pushing it
// toward zero. This creates unpredictable deflection angles after paddle hits.
func (v *Vector) Randomize() {
	perturbation := rand.Float64()
	if v.Y > 0 {
		v.Y -= perturbation
	} else {
		v.Y += perturbation
	}
}

// det returns the determinant of the 2x2 matrix formed by vectors a and b.
// Used for line intersection calculations.
func det(a, b Vector) float64 {
	return a.X*b.Y - a.Y*b.X
}
