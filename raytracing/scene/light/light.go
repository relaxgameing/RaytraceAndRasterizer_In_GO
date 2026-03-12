package light

import (
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/entity"
)

const (
	Epsilon float32 = 0.001
)

type LightType int

const (
	POINT_LIGHT LightType = iota
	DIRECTIONAL_LIGHT
	AMBIENT_LIGHT
)

type Light interface {
	GetType() LightType
	IsPointInFov(pointToCheck geom.WorldPoint, sceneEntities []entity.Entity) bool
	ComputeDiffuseReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector) float32
	ComputeSpecularReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector, specular float32, cameraPosition geom.WorldPoint) float32
}
