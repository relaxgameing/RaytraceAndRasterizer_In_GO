package viewfrustum

import (
	"math"

	"github.com/relaxgameing/computerGraphics/geom"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

const (
	BySqrt2 = 1 / math.Sqrt2
)

type InteractionWithPlane int

const (
	InsideVolume InteractionWithPlane = iota
	PartiallyInside
	OutsideVolume
)

type ViewFrustum interface {
	ObjectInsideFrustum(origin homocoord.Vec3, radius float32) InteractionWithPlane
	TriangleInsideFrustum(triangle geom.Triangle) []geom.Triangle
}
