package geom

import (
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
	"github.com/veandco/go-sdl2/sdl"
)

type Triangle struct {
	a homocoord.Vec4
	b homocoord.Vec4
	c homocoord.Vec4

	color sdl.Color
}

func NewTriangle(a, b, c homocoord.Vec4, color sdl.Color) *Triangle {
	triangle := &Triangle{a, b, c, color}

	return triangle
}
