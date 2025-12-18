package geom

import "math"

type Point struct {
	X int
	Y int
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
	if p1.Y <= p2.Y {
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
