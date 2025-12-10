package light

import (
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/veandco/go-sdl2/sdl"
)

type DirectionalLight struct {
	style     LightType
	direction geom.Vector
	intensity float32
	color     sdl.Color
}

func NewDirectionalLight(direction geom.Vector, intensity float32, color sdl.Color) *DirectionalLight {
	return &DirectionalLight{DIRECTIONAL_LIGHT, direction, intensity, color}
}

func (d *DirectionalLight) GetType() LightType {
	return d.style
}

func (d *DirectionalLight) ComputeLightingIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector) float32 {
	dot := geom.DotProduct(normalVectorOfPoint, d.direction)
	if dot <= 0 {
		return 0
	}

	return d.intensity * (dot /
		(normalVectorOfPoint.Magnitude() * d.direction.Magnitude()))
}
