package geom

import (
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
	"github.com/veandco/go-sdl2/sdl"
)

type Triangle struct {
	a homocoord.Vec3
	b homocoord.Vec3
	c homocoord.Vec3

	color sdl.Color
}

func NewTriangle(a, b, c homocoord.Vec3, color sdl.Color) *Triangle {
	triangle := &Triangle{a, b, c, color}

	return triangle
}

func (t *Triangle) GetVertex(i int) homocoord.Vec3 {
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
