package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Triangle struct {
	a WorldPoint
	b WorldPoint
	c WorldPoint

	color sdl.Color
}

func NewTriangle(a, b, c WorldPoint, color sdl.Color) *Triangle {
	triangle := &Triangle{a, b, c, color}

	return triangle
}
