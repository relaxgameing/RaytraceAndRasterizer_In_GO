package entity

import (
	"github.com/google/uuid"
	"github.com/relaxgameing/computerGraphics/raytracing/geom"
	"github.com/veandco/go-sdl2/sdl"
)

type Entity interface {
	GetId() uuid.UUID
	GetColor() sdl.Color
	GetOrigin() geom.WorldPoint
	GetSpecularExponent() float32
	GetReflectiveCoefficient() float32
	IsRayIntersecting(ray geom.Ray) (t float32, hit bool)
}
