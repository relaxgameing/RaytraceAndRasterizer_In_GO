package light

import (
	"math"

	"github.com/relaxgameing/computerGraphics/raytracing/geom"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/entity"
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
func (d *DirectionalLight) IsPointInFov(pointToCheck geom.WorldPoint, sceneEntities []entity.Entity) bool {

	lightRay := geom.Ray{
		Point:           pointToCheck,
		Lambda:          1,
		DirectionVector: d.direction,
	}

	for _, entity := range sceneEntities {
		if t, hit := entity.IsRayIntersecting(lightRay); hit && t > Epsilon {
			return false
		}
	}

	return true
}

func (d *DirectionalLight) ComputeDiffuseReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector) float32 {
	dot := geom.DotProduct(normalVectorOfPoint, d.direction)
	if dot <= 0 {
		return 0
	}

	return d.intensity * (dot /
		(normalVectorOfPoint.Magnitude() * d.direction.Magnitude()))
}

func (d *DirectionalLight) ComputeSpecularReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector, specular float32, cameraPosition geom.WorldPoint) float32 {

	dot := geom.DotProduct(normalVectorOfPoint, d.direction)

	reflectedVector := geom.NewVector(normalVectorOfPoint.ScalarProduct(2*dot).WorldPoint, d.direction.WorldPoint)
	viewVector := geom.NewVector(cameraPosition, point)

	dot = reflectedVector.Dot(*viewVector)

	if dot <= 0 {
		return 0
	}

	return float32(math.Pow(float64(dot/(reflectedVector.Magnitude()*viewVector.Magnitude())), float64(specular)))

}
