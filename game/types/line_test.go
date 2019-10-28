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
		Line{Start: Vector{0, 0}, End: Vector{10, 10}},
		Line{Start: Vector{0, 10}, End: Vector{10, 0}},
		Vector{X: 5, Y: 5},
		nil},
	{"Lines intersect in inverse direction",
		Line{Start: Vector{20, 20}, End: Vector{0, 0}},
		Line{Start: Vector{0, 20}, End: Vector{20, 0}},
		Vector{X: 10, Y: 10},
		nil},
	{"No intersection",
		Line{Start: Vector{0, 0}, End: Vector{0, 1}},
		Line{Start: Vector{1, 0}, End: Vector{1, 1}},
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
