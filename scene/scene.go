package scene

import (
	"github.com/relaxgameing/computerGraphics/scene/entity"
)

type Scene struct {
	Name string

	Canvas
	ViewPort

	SceneEntities []entity.Entity
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
			Width:              2,
			Height:             2,
			DistanceFromOrigin: 1,
		},
	}
}

func (s *Scene) AddEntity(obj ...entity.Entity) {
	s.SceneEntities = append(s.SceneEntities, obj...)
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
