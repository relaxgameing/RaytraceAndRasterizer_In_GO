package parser

import (
	"github.com/relaxgameing/computerGraphics/geom"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

type Model struct {
	name      string
	vertices  []homocoord.Vec4
	triangles []geom.Triangle
}

type ModelInstance struct {
	name        string
	model       *Model
	scale       homocoord.Vec3
	translation homocoord.Vec3
	rotation    homocoord.Mat4
}

func NewModel(name string, v []homocoord.Vec4, t []geom.Triangle) *Model {
	return &Model{
		vertices:  v,
		triangles: t,
	}
}

func EmptyModel() *Model {
	return &Model{}
}

func (m *Model) AddVertices(v ...homocoord.Vec4) {
	m.vertices = append(m.vertices, v...)
}

func (m *Model) AddTriangles(t ...geom.Triangle) {
	m.triangles = append(m.triangles, t...)
}
