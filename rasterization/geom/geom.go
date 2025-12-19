package geom

func Swap[T ~int | ~float32 | ~float64](x, y T) (T, T) {
	return y, x
}

// A + (B - A)T
func Interpolate[T ~int | ~float32 | ~float64](start, end T, step float32) float32 {
	s, e := float32(start), float32(end)
	return s + (e-s)*step
}

// * From lower value to larger value
func InterpolateAlongLine[T ~int | ~float32 | ~float64](i1, d1, i2, d2 T) []float32 {
	slope := float32(d1-d2) / float32(i1-i2)
	prev := float32(d1)

	points := make([]float32, 0)
	for i := i1; i < i2; i++ {
		points = append(points, prev)
		prev = prev + slope
	}

	return points
}
