package entity

import (
	"math"

	"github.com/google/uuid"
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/veandco/go-sdl2/sdl"
)

type Sphere struct {
	id             uuid.UUID
	origin         geom.WorldPoint
	radius         float32
	color          sdl.Color
	specular       float32
	reflectiveness float32
}

func NewSphere(origin geom.WorldPoint, radius float32, color sdl.Color, specular float32, reflectiveness float32) *Sphere {
	id, _ := uuid.NewUUID()
	return &Sphere{
		id:             id,
		origin:         origin,
		radius:         radius,
		specular:       specular,
		color:          color,
		reflectiveness: reflectiveness,
	}
}

func (s *Sphere) GetOrigin() geom.WorldPoint {
	return s.origin
}

func (s *Sphere) GetId() uuid.UUID {
	return s.id
}

func (s *Sphere) GetColor() sdl.Color {
	return s.color
}

func (s *Sphere) GetSpecularExponent() float32 {
	return s.specular
}

func (s *Sphere) GetReflectiveCoefficient() float32 {
	return s.reflectiveness
}

func (s *Sphere) IsRayIntersecting(ray geom.Ray) (t float32, hit bool) {
	co := geom.NewVector(ray.Point, s.origin)

	a := geom.DotProduct(ray.DirectionVector, ray.DirectionVector)
	b := 2 * geom.DotProduct(*co, ray.DirectionVector)
	c := geom.DotProduct(*co, *co) - s.radius*s.radius

	discriminant := float64(b*b - 4*a*c)
	if discriminant < 0 {
		return math.SmallestNonzeroFloat32, false
	}

	t1 := (-b + float32(math.Sqrt(discriminant))) / (2 * a)
	t2 := (-b - float32(math.Sqrt(discriminant))) / (2 * a)

	small, big := min(t1, t2), max(t1, t2)

	if small >= 0 {
		return small, true
	}

	if big >= 0 {
		return big, true
	}

	return math.SmallestNonzeroFloat32, false

}
