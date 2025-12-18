package geom

type Line struct {
	//* NOTE: start and end are just name there is no ordering
	start Point
	end   Point

	intercept float32
	slope     float32
}

func NewLine(p1, p2 Point) *Line {
	slope := CalculateSlope(p1, p2)
	return &Line{
		start:     p1,
		end:       p2,
		slope:     slope,
		intercept: CalculateIntercept(p1, slope),
	}
}

func CalculateSlope(p1, p2 Point) float32 {
	return (p1.Y - p2.Y) / (p1.X - p2.X)
}

func CalculateIntercept(p1 Point, slope float32) float32 {
	return p1.Y - slope*p1.X
}
