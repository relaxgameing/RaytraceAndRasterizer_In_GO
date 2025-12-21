package parser

import (
	"github.com/relaxgameing/computerGraphics/geom"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

type Model struct {
	vertices []homocoord.Vec4

	triangles []geom.Triangle
}

type ModelInstance struct {
	model *Model
}

func NewModel(v []homocoord.Vec4, t []geom.Triangle) *Model {
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
