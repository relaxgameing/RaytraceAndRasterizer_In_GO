package geom

import (
	"math"
)

type Vector struct {
	WorldPoint
}

type WorldPoint struct {
	X float32
	Y float32
	Z float32
}

func (w WorldPoint) ToVector() *Vector {
	return &Vector{
		WorldPoint: w,
	}
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

func (v *Vector) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y) + float64(v.Z*v.Z)))
}

func (v Vector) ScalarProduct(factor float32) *Vector {
	return &Vector{
		WorldPoint{
			v.X * factor,
			v.Y * factor,
			v.Z * factor,
		},
	}
}

func (v *Vector) Add(a Vector) *Vector {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
	return v
}

func (v *Vector) Dot(b Vector) float32 {
	return v.X*b.X + v.Y*b.Y + v.Z*b.Z
}

func DotProduct(a Vector, b Vector) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (v *Vector) MirrorReflectVector(vectorToReflect Vector) *Vector {
	return NewVector(v.ScalarProduct(2*v.Dot(vectorToReflect)).WorldPoint, vectorToReflect.WorldPoint)
}

func (v *Vector) UnitVector() *Vector {
	return v.ScalarProduct(1 / v.Magnitude())
}

// * Rotates the vector By the given rotation
// ? NOTE: Modifies Caller Instance
func (v *Vector) Rotate(rotation Rotation) *Vector {
	v.ChangePitch(float64(rotation.Pitch)).
		ChangeRoll(float64(rotation.Roll)).
		ChangeYaw(float64(rotation.Yaw))
	return v
}

func (v *Vector) ChangeYaw(degree float64) *Vector {
	radian := DegreeToRadian(degree)
	cosVal := float32(math.Cos(float64(radian)))
	sinVal := float32(math.Sin(float64(radian)))
	newX := v.X*cosVal + v.Z*sinVal
	newZ := -v.X*sinVal + v.Z*cosVal
	v.X = newX
	v.Z = newZ
	return v
}

func (v *Vector) ChangePitch(degree float64) *Vector {
	radian := DegreeToRadian(degree)
	cosVal := float32(math.Cos(float64(radian)))
	sinVal := float32(math.Sin(float64(radian)))
	newY := v.Y*cosVal - v.Z*sinVal
	newZ := v.Y*sinVal + v.Z*cosVal
	v.Y = newY
	v.Z = newZ
	return v
}

func (v *Vector) ChangeRoll(degree float64) *Vector {
	radian := DegreeToRadian(degree)
	cosVal := float32(math.Cos(float64(radian)))
	sinVal := float32(math.Sin(float64(radian)))
	newX := v.X*cosVal + (-1)*v.Y*sinVal
	newY := v.X*sinVal + v.Y*cosVal
	v.X = newX
	v.Y = newY
	return v
}
