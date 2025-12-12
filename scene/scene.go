package scene

import (
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/scene/entity"
	"github.com/relaxgameing/computerGraphics/scene/light"
)

type Scene struct {
	Name string

	Canvas
	ViewPort
	MainCamera *Camera

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
			DistanceFromCamera: 1,
		},
		MainCamera: &Camera{
			Position: geom.WorldPoint{X: 0, Y: 0, Z: 0},
			Rotation: geom.Rotation{Pitch: 0, Yaw: 0, Roll: 0},
		},
	}
}

func (s *Scene) AddEntity(obj ...entity.Entity) {
	s.SceneEntities = append(s.SceneEntities, obj...)
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
	DistanceFromCamera float32
}

/*
*	Camera:
*	it is our eye through which we can see the scene
 */
type Camera struct {
	//* this is the position of the camera in the scene
	Position geom.WorldPoint

	//* ViewDirection is a unit vector in the direction where the camera is seeing
	Rotation geom.Rotation
}
