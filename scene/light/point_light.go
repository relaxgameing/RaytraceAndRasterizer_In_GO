package light

import (
	"math"

	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/scene/entity"
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

func (p *PointLight) IsPointInFov(pointToCheck geom.WorldPoint, sceneEntities []entity.Entity) bool {

	lightRay := geom.Ray{
		Point:           pointToCheck,
		Lambda:          1,
		DirectionVector: *geom.NewVector(p.source, pointToCheck),
	}

	for _, entity := range sceneEntities {
		t, hit := entity.IsRayIntersecting(lightRay)
		if hit && Epsilon < t && t <= 1 {
			return false
		}
	}

	return true
}

func (p *PointLight) ComputeDiffuseReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector) float32 {
	lightVector := geom.NewVector(p.source, point)
	dot := geom.DotProduct(normalVectorOfPoint, *lightVector)
	if dot <= 0 {
		return 0
	}

	return p.intensity * (dot /
		(normalVectorOfPoint.Magnitude() * lightVector.Magnitude()))
}

func (p *PointLight) ComputeSpecularReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector, specular float32, cameraPosition geom.WorldPoint) float32 {

	lightVector := geom.NewVector(p.source, point)
	dot := geom.DotProduct(normalVectorOfPoint, *lightVector)

	reflectedVector := geom.NewVector(normalVectorOfPoint.ScalarProduct(2*dot).WorldPoint, lightVector.WorldPoint)
	viewVector := geom.NewVector(cameraPosition, point)

	dot = reflectedVector.Dot(*viewVector)

	if dot <= 0 {
		return 0
	}

	return float32(math.Pow(float64(dot/(reflectedVector.Magnitude()*viewVector.Magnitude())), float64(specular)))

}
