package game

type Vector struct {
	x, y int
}

func Add(v1 Vector, v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}

func Up() Vector {
	return Vector{0, -1}
}

func Right() Vector {
	return Vector{-1, 0}
}

func Down() Vector {
	return Vector{0, 1}
}

func Left() Vector {
	return Vector{1, 0}
}
