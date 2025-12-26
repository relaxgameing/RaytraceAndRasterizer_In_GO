package scene

import (
	"github.com/relaxgameing/computerGraphics/editor/scene"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
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
			ViewCamera: scene.NewCamera(homocoord.Vec3{0, 0, 0},
				homocoord.Vec3{0, 0, 1},
				homocoord.Vec3{0, 1, 0},
				homocoord.IdentityMat4(),
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
