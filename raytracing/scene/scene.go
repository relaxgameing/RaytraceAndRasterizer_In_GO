package scene

import (
	"github.com/relaxgameing/computerGraphics/editor/scene"
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/entity"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/light"
)

type RayScene struct {
	scene.BaseScene

	SceneEntities []entity.Entity
	Lightings     []light.Light
}

// Todo: Dependency Injection for configuration of RayScene
func NewScene(sceneName string) *RayScene {
	return &RayScene{
		BaseScene: scene.BaseScene{
			Name: sceneName,
			Canvas: scene.Canvas{
				Width:  800,
				Height: 600,
			},
			ViewPort: scene.ViewPort{
				Width:              1,
				Height:             1,
				DistanceFromOrigin: 1,
			},
			ViewCamera: scene.NewCamera(
				geom.WorldPoint{X: 0, Y: 0, Z: 0},
				scene.InitialCameraDirection,
				geom.Rotation{Pitch: 0, Yaw: 0, Roll: 0},
			),
		},
	}
}

func (s *RayScene) SetSceneName(name string) {
	s.Name = name
}

func (s *RayScene) GetCanvasWidth() int {
	return s.Canvas.Width
}

func (s *RayScene) GetCanvasHeight() int {
	return s.Canvas.Height
}
func (s *RayScene) AddSceneEntities(entities ...entity.Entity) {
	s.SceneEntities = append(s.SceneEntities, entities...)
}

func (s *RayScene) AddLighting(lights ...light.Light) {
	s.Lightings = append(s.Lightings, lights...)
}
