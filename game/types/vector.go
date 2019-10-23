package types

type Vector struct {
	X, Y float64
}

func (v *Vector) Add(other Vector) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vector) InvertX() {
	v.X *= -1
}

func (v *Vector) InvertY() {
	v.Y *= -1
}
