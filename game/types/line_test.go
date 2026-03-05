package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var intersectTests = []struct {
	name         string
	line1        Line
	line2        Line
	intersection Vector
	err          error
}{
	{"Lines intersect",
		Line{Start: Vector{X: 0, Y: 0}, End: Vector{X: 10, Y: 10}},
		Line{Start: Vector{X: 0, Y: 10}, End: Vector{X: 10, Y: 0}},
		Vector{X: 5, Y: 5},
		nil},
	{"Lines intersect in inverse direction",
		Line{Start: Vector{X: 20, Y: 20}, End: Vector{X: 0, Y: 0}},
		Line{Start: Vector{X: 0, Y: 20}, End: Vector{X: 20, Y: 0}},
		Vector{X: 10, Y: 10},
		nil},
	{"No intersection",
		Line{Start: Vector{X: 0, Y: 0}, End: Vector{X: 0, Y: 1}},
		Line{Start: Vector{X: 1, Y: 0}, End: Vector{X: 1, Y: 1}},
		Vector{},
		fmt.Errorf("No intersection")},
}

func TestLineIntersections(t *testing.T) {
	for _, test := range intersectTests {
		t.Run(test.name, func(t *testing.T) {
			v, e := Intersect(test.line1, test.line2)
			assert.Equal(t, test.err, e)
			assert.Equal(t, test.intersection, v)
		})
	}
}
