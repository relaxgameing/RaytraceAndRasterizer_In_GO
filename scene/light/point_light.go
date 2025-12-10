package light

import (
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/veandco/go-sdl2/sdl"
)

type PointLight struct {
	style     LightType
	source    geom.WorldPoint
	intensity float32
	color     sdl.Color
}

func NewPointLight(source geom.WorldPoint, intensity float32, color sdl.Color) *PointLight {
	return &PointLight{POINT_LIGHT, source, intensity, color}
}

func (p *PointLight) GetType() LightType {
	return p.style
}
func (p *PointLight) ComputeLightingIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector) float32 {
	lightVector := geom.NewVector(p.source, point)
	dot := geom.DotProduct(normalVectorOfPoint, *lightVector)
	if dot <= 0 {
		return 0
	}

	return p.intensity * (dot /
		(normalVectorOfPoint.Magnitude() * lightVector.Magnitude()))
}
