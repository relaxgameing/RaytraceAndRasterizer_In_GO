package scene

import (
	"github.com/relaxgameing/computerGraphics/geom"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

type Model struct {
	name      string
	vertices  []homocoord.Vec3
	triangles []geom.Triangle
}

func NewModel(name string, v []homocoord.Vec3, t []geom.Triangle) *Model {
	return &Model{
		name:      name,
		vertices:  v,
		triangles: t,
	}
}

func (m Model) Name() string                   { return m.name }
func (m Model) VertexCount() int               { return len(m.vertices) }
func (m Model) TriangleCount() int             { return len(m.triangles) }
func (m Model) VertexAt(i int) homocoord.Vec3  { return m.vertices[i] }
func (m Model) TriangleAt(i int) geom.Triangle { return m.triangles[i] }

func EmptyModel() *Model {
	return &Model{}
}

type ModelInstance struct {
	name        string
	model       *Model
	scale       homocoord.Vec3
	translation homocoord.Vec3
	rotation    homocoord.Mat4
}

func NewModelInstance(name string, model *Model, scale, translation homocoord.Vec3, rotation homocoord.Mat4) ModelInstance {
	return ModelInstance{
		name:        name,
		model:       model,
		scale:       scale,
		translation: translation,
		rotation:    rotation,
	}
}

func (mi ModelInstance) Name() string {
	return mi.name
}

func (mi ModelInstance) Model() *Model {
	return mi.model
}

func (mi ModelInstance) GetScale() homocoord.Vec3 {
	return mi.scale
}

func (mi ModelInstance) GetTranslation() homocoord.Vec3 {
	return mi.translation
}

func (mi ModelInstance) GetRotation() homocoord.Mat4 {
	return mi.rotation
}
