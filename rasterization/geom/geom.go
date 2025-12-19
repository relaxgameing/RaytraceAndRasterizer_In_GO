package geom

import "math"

// A + (B - A)T
func Interpolate(start, end int, step float32) float32 {
	return float32(start) + float32(end-start)*step
}

func InterpolateAlongLine(i1, d1, i2, d2 int) []Point {
	slope := float32(d1-d2) / float32(i1-i2)
	prev := float32(d1)

	points := make([]Point, 0)
	for i := i1; i <= i2; i++ {
		points = append(points, Point{
			X: i,
			Y: int(math.Round(float64(prev))),
		})
		prev = prev + slope
	}

	return points
}
