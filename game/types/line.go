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
	xDiff := Vector{xDiff1, xDiff2}
	yDiff := Vector{yDiff1, yDiff2}
	div := det(xDiff, yDiff)
	if div == 0 {
		return Vector{}, fmt.Errorf("No intersection")
	}
	d := Vector{det(line1.Start, line1.End), det(line2.Start, line2.End)}
	return Vector{det(d, xDiff) / div, det(d, yDiff) / div}, nil
}
