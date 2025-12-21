package parser

import "github.com/relaxgameing/computerGraphics/geom"

type Model struct {
	vertices []geom.WorldPoint

	triangles []geom.Triangle
}

func NewModel(v []geom.WorldPoint, t []geom.Triangle) *Model {
	return &Model{
		vertices:  v,
		triangles: t,
	}
}

func EmptyModel() *Model {
	return &Model{}
}

func (m *Model) AddVertices(v ...geom.WorldPoint) {
	m.vertices = append(m.vertices, v...)
}

func (m *Model) AddTriangles(t ...geom.Triangle) {
	m.triangles = append(m.triangles, t...)
}
