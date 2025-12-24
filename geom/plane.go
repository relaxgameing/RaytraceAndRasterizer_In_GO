package geom

import homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"

type Plane struct {
	normal   homocoord.Vec3
	distance float32
}

func NewPlane(normal homocoord.Vec3, distanceFromOrigin float32) *Plane {
	return &Plane{
		normal:   normal,
		distance: distanceFromOrigin,
	}
}

func (p *Plane) DistanceFromPlane(point homocoord.Vec3) float32 {
	return p.normal.Dot(point) + p.distance
}

func (p *Plane) ContainsPoint(point homocoord.Vec3) bool {
	if p.DistanceFromPlane(point) == 0 {
		return true
	}
	return false
}

// * returns point of intersection else false if not point of intersection
func (p *Plane) IntersectsLine(startPoint, endpoint homocoord.Vec3) (homocoord.Vec3, bool) {
	distA := p.DistanceFromPlane(startPoint)
	distB := p.DistanceFromPlane(endpoint)

	var insideVolPoint, outsideVolPoint homocoord.Vec3 = startPoint, endpoint

	if max(distA, distB) > 0 && min(distA, distB) < 0 {
		if distA < distB {
			insideVolPoint, outsideVolPoint = endpoint, startPoint
		}
		t := (-p.distance - p.normal.Dot(insideVolPoint)) / (p.normal.Dot(outsideVolPoint.Subtract(insideVolPoint)))
		return LerpOnLine(insideVolPoint, outsideVolPoint, t), true
	}

	return homocoord.Vec3{}, false
}
