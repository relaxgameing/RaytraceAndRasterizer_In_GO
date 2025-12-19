package shape

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/rasterization/geom"
	"github.com/veandco/go-sdl2/sdl"
)

type Triangle struct {
	a geom.Point
	b geom.Point
	c geom.Point

	color sdl.Color
}

type TriangleOptions func(t *Triangle)

func NewTriangle(a, b, c geom.Point, options ...TriangleOptions) *Triangle {
	triangle := &Triangle{a, b, c, common.ColorRed}

	for _, op := range options {
		op(triangle)
	}

	return triangle
}

// * Triangle Options
func (t *Triangle) WithColor(color sdl.Color) *Triangle {
	t.color = color
	return t
}

//* Shape Interface

func (t *Triangle) GetColor() sdl.Color {
	return t.color
}

func (t *Triangle) Draw() []*geom.Point {
	trianglePoints := make([]*geom.Point, 0)

	// Outline
	trianglePoints = append(trianglePoints, NewLine(t.a, t.b, WithColor(t.color)).Draw()...)
	trianglePoints = append(trianglePoints, NewLine(t.a, t.c, WithColor(t.color)).Draw()...)
	trianglePoints = append(trianglePoints, NewLine(t.c, t.b, WithColor(t.color)).Draw()...)
	log.Info("Triangle", "len of outline", len(trianglePoints))

	trianglePoints = append(trianglePoints, t.fillTriangle()...)
	log.Info("Triangle", "len of filled triangle", len(trianglePoints))
	return trianglePoints
}

func (t *Triangle) fillTriangle() []*geom.Point {
	points := make([]*geom.Point, 0)

	x, y := geom.UpperPoint(t.a, t.b)
	top, z := geom.UpperPoint(x, t.c)
	mid, bottom := geom.UpperPoint(y, z)

	topMidSide := NewLine(top, mid)
	topBottomSide := NewLine(top, bottom)
	// top -> mid
	for i := top.Y; i >= mid.Y; i-- {
		midSidePoint := geom.Point{X: int(topMidSide.ComputeXForY(i)), Y: i}
		bottomSidePoint := geom.Point{X: int(topBottomSide.ComputeXForY(i)), Y: i}
		points = append(points, NewLine(midSidePoint, bottomSidePoint).Draw()...)
	}

	// mid -> bottom
	midBottomSide := NewLine(mid, bottom)
	for i := mid.Y; i >= bottom.Y; i-- {
		topBottomSidePoint := geom.NewPoint(int(topBottomSide.ComputeXForY(i)), i)
		midBottomSidePoint := geom.NewPoint(int(midBottomSide.ComputeXForY(i)), i)
		points = append(points, NewLine(*topBottomSidePoint, *midBottomSidePoint).Draw()...)
	}

	return points
}
