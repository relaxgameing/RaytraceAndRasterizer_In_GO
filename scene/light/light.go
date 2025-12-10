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
	ComputeLightingIntensityOfPoint(point geom.WorldPoint, normalOfPoint geom.Vector) float32
}
