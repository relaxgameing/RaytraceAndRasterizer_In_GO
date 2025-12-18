package scene

import (
	"github.com/relaxgameing/computerGraphics/raytracing/scene/entity"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/light"
)

type Scene struct {
	Name string

	Canvas
	ViewPort

	SceneEntities []entity.Entity
	Lightings     []light.Light
}

// Todo: Dependency Injection for configuration of Scene
func NewScene(sceneName string) *Scene {
	return &Scene{
		Name: sceneName,
		Canvas: Canvas{
			Width:  800,
			Height: 600,
		},
		ViewPort: ViewPort{
			Width:              1,
			Height:             1,
			DistanceFromOrigin: 1,
		},
	}
}

func (s *Scene) SetSceneName(name string) {
	s.Name = name
}

func (s *Scene) GetWidth() int {
	return s.Canvas.Width
}

func (s *Scene) GetHeight() int {
	return s.Canvas.Height
}
func (s *Scene) AddSceneEntities(entities ...entity.Entity) {
	s.SceneEntities = append(s.SceneEntities, entities...)
}

func (s *Scene) AddLighting(lights ...light.Light) {
	s.Lightings = append(s.Lightings, lights...)
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
