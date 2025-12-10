package light

import (
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/veandco/go-sdl2/sdl"
)

type AmbientLight struct {
	style     LightType
	intensity float32
	color     sdl.Color
}

func NewAmbientLight(intensity float32, color sdl.Color) *AmbientLight {
	return &AmbientLight{AMBIENT_LIGHT, intensity, color}
}

func (a *AmbientLight) GetType() LightType {
	return a.style
}

func (a *AmbientLight) ComputeLightingIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector) float32 {
	return a.intensity
}
