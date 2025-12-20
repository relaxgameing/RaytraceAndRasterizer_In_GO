package light

import (
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/entity"
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

func (a *AmbientLight) IsPointInFov(pointToCheck geom.WorldPoint, sceneEntities []entity.Entity) bool {
	return true
}

func (a *AmbientLight) ComputeDiffuseReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector) float32 {
	return a.intensity
}

func (a *AmbientLight) ComputeSpecularReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector, specular float32, cameraPosition geom.WorldPoint) float32 {
	return 0
}
