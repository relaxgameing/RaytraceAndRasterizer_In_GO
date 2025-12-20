package scene

import "github.com/relaxgameing/computerGraphics/rasterization/scene/shape"

type Scene struct {
	name string
	Canvas
	ViewPort

	shapes []shape.Shape
}

/*
* Canvas:
* it is the screen which we are able to see in the compute
* it's unit is pixels
* it is a 2D canvas
 */

type Canvas struct {
	Width  int
	Height int
}

/*
*ViewPort:
* it is the window through which we see the real world
* it is world units
* it is a 3D world
 */

type ViewPort struct {
	Width              int
	Height             int
	DistanceFromOrigin int
}

type sceneOptions func(s *Scene)

func NewScene(opt ...sceneOptions) *Scene {
	scene := &Scene{
		name: "Rasterization",
		Canvas: Canvas{Width: 800,
			Height: 600},
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
		s.Canvas.Width = width
	}
}
func WithHeight(height int) sceneOptions {
	return func(s *Scene) {
		s.Canvas.Height = height
	}
}

// * Scene interface
func (s *Scene) SetSceneName(name string) {
	s.name = name
}

func (s *Scene) GetCanvasWidth() int {
	return s.Canvas.Width
}

func (s *Scene) GetCanvasHeight() int {
	return s.Canvas.Height
}

func (s *Scene) GetShapes() []shape.Shape {
	return s.shapes
}

func (s *Scene) AddSceneEntities(entities ...shape.Shape) {
	s.shapes = append(s.shapes, entities...)
}

func (s *Scene) CanvasToSdl(cx, cy int) (x, y int) {
	x = (s.Canvas.Width / 2) + cx
	y = (s.Canvas.Height / 2) - cy

	return x, y
}
