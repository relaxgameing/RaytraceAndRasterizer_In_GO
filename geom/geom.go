package geom

import (
	"math"

	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

func Abs[T ~int | ~int32 | ~float32 | ~float64](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

// pi radian = 180 deg
func DegreeToRadian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func RadianToDegree(radian float64) float64 {
	return radian * (180 / math.Pi)
}

func Swap[T ~int | ~float32 | ~float64](x, y T) (T, T) {
	return y, x
}

// A + (B - A)T
func Lerp[T ~int | ~float32 | ~float64](start, end T, step float32) float32 {
	s, e := float32(start), float32(end)
	return s + (e-s)*step
}

func LerpOnLine(start, end homocoord.Vec3, t float32) homocoord.Vec3 {
	diff := end.Subtract(start)
	return start.Add(diff.ScalarPrd(t))
}

// * From lower value to larger value
func InterpolateAlongLine[T ~int | ~float32 | ~float64](i1, d1, i2, d2 T) []float32 {
	slope := float32(d1-d2) / float32(i1-i2)
	prev := float32(d1)

	points := make([]float32, 0)
	for i := i1; i <= i2; i++ {
		points = append(points, prev)
		prev = prev + slope
	}

	return points
}
