package geom

import (
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
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

	idx := 0
	for i := bottom.Y + 1; i <= top.Y; i++ {
		longSidePoint := LerpOnLine(bottom.Vec3, top.Vec3,
			(float32(idx))/Abs(top.Y-bottom.Y-1))

		var otherSidePoint homocoord.Vec3
		if i > (mid.Y) {
			otherSidePoint = LerpOnLine(mid.Vec3, top.Vec3,
				(float32(idx)-Abs(mid.Y-bottom.Y))/Abs(top.Y-mid.Y-1),
			)
		} else {
			otherSidePoint = LerpOnLine(
				bottom.Vec3, mid.Vec3,
				float32(idx)/Abs(mid.Y-bottom.Y-1),
			)
		}

		points = append(points, NewLine(
			*NewPointFromVec3(longSidePoint),
			*NewPointFromVec3(otherSidePoint),
		).Draw()...)
		idx++
	}

	return points
}
