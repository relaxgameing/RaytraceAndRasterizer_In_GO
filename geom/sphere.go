package geom

import homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"

type Sphere struct {
	origin homocoord.Vec3
	radius float32
}

func NewSphere(origin homocoord.Vec3, radius float32) *Sphere {
	return &Sphere{
		origin: origin,
		radius: radius,
	}
}

func (s *Sphere) GetOrigin() homocoord.Vec3 {
	return s.origin
}

func (s *Sphere) GetRadius() float32 {
	return s.radius
}
