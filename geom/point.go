package geom

import (
	"math"

	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

type Point struct {
	homocoord.Vec3
	Intensity float32
}

type PointOption func(*Point)

func NewPoint(x, y, z float32, options ...PointOption) *Point {
	p := &Point{
		homocoord.Vec3{
			X: (x),
			Y: (y),
			Z: z,
		},
		1,
	}

	for _, op := range options {
		op(p)
	}

	return p
}

func NewPointFromVec3(v homocoord.Vec3, options ...PointOption) *Point {
	p := &Point{
		homocoord.Vec3{
			X: v.X,
			Y: v.Y,
			Z: v.Z,
		},
		1,
	}

	for _, op := range options {
		op(p)
	}

	return p
}

func PointWithIntensity(intensity float32) PointOption {
	return func(p *Point) {
		p.Intensity = intensity
	}
}

func DistanceFromOrigin(p Point) float32 {
	return float32(math.Sqrt(float64(p.X*p.X) + float64(p.Y*p.Y)))
}

// * returns the point more towards the left
func LeftPoint(p1, p2 Point) (left, right Point) {
	if p1.X < p2.X {
		return p1, p2
	}
	return p2, p1
}

// *returns the point more towards the top
func UpperPoint(p1, p2 Point) (upper, lower Point) {
	if p1.Y >= p2.Y {
		return p1, p2
	}
	return p2, p1
}

func PointCloserToOrigin(p1, p2 Point) (start, end Point) {
	if DistanceFromOrigin(p1) < DistanceFromOrigin(p2) {
		return p1, p2
	}
	return p2, p1
}
