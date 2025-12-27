package shape

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/geom"
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
	// trianglePoints = append(trianglePoints, t.fillTriangle()...)
	log.Info("Triangle", "len of filled triangle", len(trianglePoints))
	return trianglePoints
}

// func (t *Triangle) fillTriangle() []*geom.Point {
// 	points := make([]*geom.Point, 0)

// 	x, y := geom.UpperPoint(t.a, t.b)
// 	top, z := geom.UpperPoint(x, t.c)
// 	mid, bottom := geom.UpperPoint(y, z)

// 	// top -> mid
// 	topMidSidePoints := geom.InterpolateAlongLine(mid.Y, mid.X, top.Y, top.X)
// 	intensitiesOfTopMid := geom.InterpolateAlongLine(float32(mid.Y), mid.Intensity, float32(top.Y), top.Intensity)

// 	topBottomSidePoints := geom.InterpolateAlongLine(bottom.Y, bottom.X, top.Y, top.X)
// 	intensitiesOfTopBottom := geom.InterpolateAlongLine(float32(bottom.Y), bottom.Intensity, float32(top.Y), top.Intensity)

// 	midBottomSidePoints := geom.InterpolateAlongLine(bottom.Y, bottom.X, mid.Y, mid.X)
// 	intensitiesOfMidBottom := geom.InterpolateAlongLine(float32(bottom.Y), bottom.Intensity, float32(mid.Y), mid.Intensity)

// 	idx := 0
// 	for i := bottom.Y + 1; i < top.Y; i++ {
// 		var longSidePoint, otherSidePoint geom.Point

// 		if i > mid.Y {
// 			otherSidePoint = *geom.NewPoint(int(topMidSidePoints[idx-(mid.Y-bottom.Y)]), i, geom.PointWithIntensity(intensitiesOfTopMid[idx-(mid.Y-bottom.Y)]))
// 		} else {
// 			otherSidePoint = *geom.NewPoint(int(midBottomSidePoints[idx]), i, geom.PointWithIntensity(intensitiesOfMidBottom[idx]))
// 		}

// 		longSidePoint = *geom.NewPoint(int(topBottomSidePoints[idx]), i, geom.PointWithIntensity(intensitiesOfTopBottom[idx]))

// 		points = append(points, NewLine(longSidePoint, otherSidePoint).Draw()...)
// 		idx++
// 	}

// 	return points
// }
