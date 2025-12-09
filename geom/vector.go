package geom

import "math"

type Vector struct {
	WorldPoint
}

type WorldPoint struct {
	X float32
	Y float32
	Z float32
}

func (v *Vector) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y) + float64(v.Z*v.Z)))
}

func NewVector(head WorldPoint, tail WorldPoint) *Vector {
	return &Vector{
		WorldPoint{
			X: head.X - tail.X,
			Y: head.Y - tail.Y,
			Z: head.Z - tail.Z,
		},
	}
}

func DotProduct(a Vector, b Vector) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}
