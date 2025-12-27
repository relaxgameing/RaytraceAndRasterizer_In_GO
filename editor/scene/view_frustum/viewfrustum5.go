package viewfrustum

import (
	"github.com/relaxgameing/computerGraphics/geom"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

type ViewFrustum5 struct {
	planes []*geom.Plane
}

func New5PlaneFrustum() *ViewFrustum5 {
	return &ViewFrustum5{
		planes: []*geom.Plane{
			geom.NewPlane(homocoord.Vec3{X: 0, Y: 0, Z: 1}, 1),
			geom.NewPlane(homocoord.Vec3{X: BySqrt2, Y: 0, Z: BySqrt2}, 0),
			geom.NewPlane(homocoord.Vec3{X: -BySqrt2, Y: 0, Z: BySqrt2}, 0),
			geom.NewPlane(homocoord.Vec3{X: 0, Y: BySqrt2, Z: BySqrt2}, 0),
			geom.NewPlane(homocoord.Vec3{X: 0, Y: -BySqrt2, Z: BySqrt2}, 0),
		},
	}
}

func (view *ViewFrustum5) ObjectInsideFrustum(origin homocoord.Vec3, radius float32) InteractionWithPlane {
	for _, plane := range view.planes {
		dist := plane.DistanceFromPlane(origin)
		if dist < 0 {
			if -dist > radius {
				return OutsideVolume
			} else {
				return PartiallyInside
			}
		}
	}

	return InsideVolume
}

func (view *ViewFrustum5) TriangleInsideFrustum(triangle geom.Triangle) []geom.Triangle {

	for _, plane := range view.planes {
		triangles, ok := view.TriangleInsidePlane(*plane, triangle)
		if !ok {
			return triangles
		}
	}

	return []geom.Triangle{triangle}
}

func (view *ViewFrustum5) TriangleInsidePlane(plane geom.Plane, triangle geom.Triangle) ([]geom.Triangle, bool) {
	inside := make([]geom.Point, 0)
	outside := make([]geom.Point, 0)
	for i := 0; i < 2; i++ {
		vertex := triangle.GetVertex(i)
		if plane.DistanceFromPlane(vertex.Vec3) > 0 {
			inside = append(inside, vertex)
		} else {
			outside = append(outside, vertex)
		}
	}

	switch len(inside) {
	case 0:
		return []geom.Triangle{}, false
	case 1:
		iPoint := inside[0]
		edgePoint1, _ := plane.IntersectsLine(iPoint.Vec3, outside[0].Vec3)
		edgePoint2, _ := plane.IntersectsLine(iPoint.Vec3, outside[1].Vec3)

		return []geom.Triangle{
			*geom.NewTriangle(iPoint, *geom.NewPointFromVec3(edgePoint1), *geom.NewPointFromVec3(edgePoint2), triangle.GetColor()),
		}, false
	case 2:
		iPoint1 := inside[0]
		iPoint2 := inside[1]
		outside := outside[0]

		x, _ := plane.IntersectsLine(iPoint1.Vec3, outside.Vec3)
		y, _ := plane.IntersectsLine(iPoint2.Vec3, outside.Vec3)
		return []geom.Triangle{
			*geom.NewTriangle(iPoint1, *geom.NewPointFromVec3(x), iPoint2, triangle.GetColor()),
			*geom.NewTriangle(*geom.NewPointFromVec3(x), *geom.NewPointFromVec3(y), iPoint2, triangle.GetColor()),
		}, false
	default:
		return []geom.Triangle{triangle}, true
	}
}
