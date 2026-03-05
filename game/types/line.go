package types

import "fmt"

type Line struct {
	Start, End Vector
}

func (l *Line) diff() (float64, float64) {
	return l.Start.X - l.End.X, l.Start.Y - l.End.Y
}

func Intersect(line1, line2 Line) (Vector, error) {
	xDiff1, yDiff1 := line1.diff()
	xDiff2, yDiff2 := line2.diff()
	xDiff := Vector{X: xDiff1, Y: xDiff2}
	yDiff := Vector{X: yDiff1, Y: yDiff2}
	div := det(xDiff, yDiff)
	if div == 0 {
		return Vector{}, fmt.Errorf("No intersection")
	}
	d := Vector{X: det(line1.Start, line1.End), Y: det(line2.Start, line2.End)}
	return Vector{X: det(d, xDiff) / div, Y: det(d, yDiff) / div}, nil
}
