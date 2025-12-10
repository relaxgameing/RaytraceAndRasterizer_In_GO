package light

import "github.com/relaxgameing/computerGraphics/geom"

type LightType int

const (
	POINT_LIGHT LightType = iota
	DIRECTIONAL_LIGHT
	AMBIENT_LIGHT
)

type Light interface {
	GetType() LightType
	ComputeDiffuseReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector) float32
	ComputeSpecularReflectionIntensityOfPoint(point geom.WorldPoint, normalVectorOfPoint geom.Vector, specular float32, cameraPosition geom.WorldPoint) float32
}
