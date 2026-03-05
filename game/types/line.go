package types

import (
	"errors"
	"math"
)

// ErrNoIntersection is returned when two lines are parallel and do not intersect.
var ErrNoIntersection = errors.New("lines are parallel and do not intersect")

// Line represents a line segment defined by a start and end point.
type Line struct {
	Start, End Vector
}

// Length returns the length of the line segment.
func (l *Line) Length() float64 {
	dx := l.End.X - l.Start.X
	dy := l.End.Y - l.Start.Y
	return math.Hypot(dx, dy)
}

// Direction returns a unit vector pointing from Start to End.
// Returns a zero vector if the line has zero length.
func (l *Line) Direction() Vector {
	dir := Vector{X: l.End.X - l.Start.X, Y: l.End.Y - l.Start.Y}
	dir.Normalize()
	return dir
}

// Intersect finds the intersection point of two infinite lines defined by
// the given line segments. Returns ErrNoIntersection if the lines are parallel.
func Intersect(line1, line2 Line) (Vector, error) {
	// Compute direction differences for each line
	dx1 := line1.Start.X - line1.End.X
	dy1 := line1.Start.Y - line1.End.Y
	dx2 := line2.Start.X - line2.End.X
	dy2 := line2.Start.Y - line2.End.Y

	// Divisor is the determinant of the direction matrix
	div := dx1*dy2 - dy1*dx2
	if div == 0 {
		return Vector{}, ErrNoIntersection
	}

	// Determinants of start/end point matrices
	d1 := line1.Start.X*line1.End.Y - line1.Start.Y*line1.End.X
	d2 := line2.Start.X*line2.End.Y - line2.Start.Y*line2.End.X

	return Vector{
		X: (d1*dx2 - d2*dx1) / div,
		Y: (d1*dy2 - d2*dy1) / div,
	}, nil
}
