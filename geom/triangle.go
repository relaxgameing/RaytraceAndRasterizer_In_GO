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

func (t *Triangle) FillTriangle(a, b, c Point) []*Point {
	points := make([]*Point, 0)

	x, y := UpperPoint(a, b)
	top, z := UpperPoint(x, c)
	mid, bottom := UpperPoint(y, z)

	// top -> mid
	// topMidSideXVal := InterpolateAlongLine(mid.Y, mid.X, top.Y, top.X)
	// topMidSideZVal := InterpolateAlongLine(mid.Y, mid.Z, top.Y, top.Z)
	// intensitiesOfTopMid := InterpolateAlongLine(float32(mid.Y), mid.Intensity, float32(top.Y), top.Intensity)

	// topBottomSideXVal := InterpolateAlongLine(bottom.Y, bottom.X, top.Y, top.X)
	// topBottomSideZVal := InterpolateAlongLine(bottom.Y, bottom.Z, top.Y, top.Z)
	// intensitiesOfTopBottom := InterpolateAlongLine(float32(bottom.Y), bottom.Intensity, float32(top.Y), top.Intensity)

	// midBottomSideXVal := InterpolateAlongLine(bottom.Y, bottom.X, mid.Y, mid.X)
	// midBottomSideZVal := InterpolateAlongLine(bottom.Y, bottom.Z, mid.Y, mid.Z)
	// intensitiesOfMidBottom := InterpolateAlongLine(float32(bottom.Y), bottom.Intensity, float32(mid.Y), mid.Intensity)

	for i := 1; i <= int(top.Y-bottom.Y-1); i++ {
		longSidePoint := LerpOnLine(bottom.Vec3, top.Vec3, float32(i)/(top.Y-bottom.Y-1))

		var otherSidePoint homocoord.Vec3
		if i+int(bottom.Y) > int(mid.Y) {
			otherSidePoint = LerpOnLine(mid.Vec3, top.Vec3, float32(i)/(top.Y-mid.Y-1))
		} else {
			otherSidePoint = LerpOnLine(bottom.Vec3, mid.Vec3, float32(i)/(mid.Y-bottom.Y-1))
		}

		points = append(points, NewLine(
			*NewPointFromVec3(longSidePoint),
			*NewPointFromVec3(otherSidePoint),
		).Draw()...)
	}

	// idx := 0
	// for i := bottom.Y + 1; i < top.Y; i++ {
	// 	var longSidePoint, otherSidePoint Point

	// 	if i > mid.Y {
	// 		otherSidePoint = *NewPoint((topMidSideXVal[idx-int(mid.Y-bottom.Y)]), i, topMidSideZVal[idx-int(mid.Y-bottom.Y)], PointWithIntensity(intensitiesOfTopMid[idx-int(mid.Y-bottom.Y)]))
	// 	} else {
	// 		otherSidePoint = *NewPoint((midBottomSideXVal[idx]), i, midBottomSideZVal[idx], PointWithIntensity(intensitiesOfMidBottom[idx]))
	// 	}

	// 	longSidePoint = *NewPoint((topBottomSideXVal[idx]), i, topBottomSideZVal[idx], PointWithIntensity(intensitiesOfTopBottom[idx]))

	// 	points = append(points, NewLine(longSidePoint, otherSidePoint).Draw()...)
	// 	idx++
	// }

	return points
}
