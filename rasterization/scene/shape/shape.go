package shape

import (
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/veandco/go-sdl2/sdl"
)

type Shape interface {
	Draw() []*geom.Point
	GetColor() sdl.Color
}
