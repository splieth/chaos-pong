package game

import "math"

type FloatVector struct {
	x, y float64
}

type IntVector struct {
	x, y int
}

func (v FloatVector) convertToInt() IntVector {
	return IntVector{
		x: int(math.Round(v.x)),
		y: int(math.Round(v.y)),
	}
}

func (v IntVector) convertToFloat() FloatVector {
	return FloatVector{
		x: float64(v.x),
		y: float64(v.y),
	}
}

func (v IntVector) Add(other IntVector) IntVector {
	return IntVector{v.x + other.x, v.y + other.y}
}

func (v FloatVector) Add(other FloatVector) FloatVector {
	return FloatVector{v.x + other.x, v.y + other.y}
}

func (v FloatVector) Normalize() FloatVector {
	norm := math.Sqrt(v.x*v.x + v.y*v.y)
	return FloatVector{v.x / norm, v.y / norm}
}

func Up() IntVector {
	return IntVector{0, -1}
}

func Right() IntVector {
	return IntVector{-1, 0}
}

func Down() IntVector {
	return IntVector{0, 1}
}

func Left() IntVector {
	return IntVector{1, 0}
}
