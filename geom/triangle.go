package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Triangle struct {
	a Point
	b Point
	c Point

	color sdl.Color
}

func NewTriangle(a, b, c Point, color sdl.Color) *Triangle {
	triangle := &Triangle{a, b, c, color}

	return triangle
}

func (t *Triangle) GetVertex(i int) Point {
	switch i % 3 {
	case 0:
		return t.a
	case 1:
		return t.b
	case 2:
		return t.c
	default:
		return t.a
	}
}

func (t *Triangle) GetColor() sdl.Color {
	return t.color
}

// * Assumption: Input Points are already Projected
func (t *Triangle) FillTriangle(a, b, c Point) []*Point {
	points := make([]*Point, 0)

	x, y := UpperPoint(a, b)
	top, z := UpperPoint(x, c)
	mid, bottom := UpperPoint(y, z)

	longSideXPoints := InterpolateAlongLine(bottom.Y, bottom.X, top.Y, top.X)
	longSideZPoints := InterpolateAlongLine(bottom.Y, bottom.Z, top.Y, top.Z)

	topMidXPoints := InterpolateAlongLine(mid.Y, mid.X, top.Y, top.X)
	topMidZPoints := InterpolateAlongLine(mid.Y, mid.Z, top.Y, top.Z)

	midBottomXPoints := InterpolateAlongLine(bottom.Y, bottom.X, mid.Y, mid.X)
	midBottomZPoints := InterpolateAlongLine(bottom.Y, bottom.Z, mid.Y, mid.Z)

	otherSideXPoints := append(midBottomXPoints, topMidXPoints...)
	otherSideZPoints := append(midBottomZPoints, topMidZPoints...)

	idx := 0
	for i := bottom.Y; i <= top.Y; i++ {
		points = append(points, NewLine(
			*NewPoint(longSideXPoints[idx], i, longSideZPoints[idx]),
			*NewPoint(otherSideXPoints[idx], i, otherSideZPoints[idx]),
		).Draw()...)
		idx++
	}

	return points
}
