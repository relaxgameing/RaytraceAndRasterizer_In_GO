package scene

import "github.com/relaxgameing/computerGraphics/rasterization/scene/shape"

type Scene struct {
	name   string
	width  int
	height int

	shapes []shape.Shape
}

type sceneOptions func(s *Scene)

func NewScene(opt ...sceneOptions) *Scene {
	scene := &Scene{
		name:   "Rasterization",
		width:  800,
		height: 600,
	}

	for _, op := range opt {
		op(scene)
	}

	return scene
}

func WithName(name string) sceneOptions {
	return func(s *Scene) {
		s.name = name
	}
}

func WithWidth(width int) sceneOptions {
	return func(s *Scene) {
		s.width = width
	}
}
func WithHeight(height int) sceneOptions {
	return func(s *Scene) {
		s.height = height
	}
}

// * Scene interface
func (s *Scene) SetSceneName(name string) {
	s.name = name
}

func (s *Scene) GetWidth() int {
	return s.width
}

func (s *Scene) GetHeight() int {
	return s.height
}

func (s *Scene) GetShapes() []shape.Shape {
	return s.shapes
}

func (s *Scene) AddSceneEntities(entities ...shape.Shape) {
	s.shapes = append(s.shapes, entities...)
}

func (s *Scene) CanvasToSdl(cx, cy int) (x, y int) {
	x = (s.width / 2) + cx
	y = (s.height / 2) - cy

	return x, y
}
