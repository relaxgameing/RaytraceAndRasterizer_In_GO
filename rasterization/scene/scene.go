package scene

import (
	"github.com/relaxgameing/computerGraphics/editor/scene"
	"github.com/relaxgameing/computerGraphics/geom"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
	"github.com/relaxgameing/computerGraphics/rasterization/scene/shape"
)

type RasterScene struct {
	scene.BaseScene
	scene.SceneObjects
	shapes []shape.Shape
}

type sceneOptions func(s *RasterScene)

func NewRasterScene(opt ...sceneOptions) *RasterScene {
	scene := &RasterScene{
		BaseScene: scene.BaseScene{
			Name: "Rasterization",
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

	for _, op := range opt {
		op(scene)
	}

	return scene
}

func WithName(name string) sceneOptions {
	return func(s *RasterScene) {
		s.BaseScene.Name = name
	}
}

func WithWidth(width int) sceneOptions {
	return func(s *RasterScene) {
		s.Canvas.Width = width
	}
}
func WithHeight(height int) sceneOptions {
	return func(s *RasterScene) {
		s.Canvas.Height = height
	}
}

func WithSceneObjects(sceneObj *scene.SceneObjects) sceneOptions {
	return func(s *RasterScene) {
		s.SceneObjects = *sceneObj
	}
}

// * RasterScene interface
func (s *RasterScene) SetSceneName(name string) {
	s.BaseScene.Name = name
}

func (s *RasterScene) GetCanvasWidth() int {
	return s.Canvas.Width
}

func (s *RasterScene) GetCanvasHeight() int {
	return s.Canvas.Height
}

func (s *RasterScene) GetShapes() []shape.Shape {
	return s.shapes
}

func (s *RasterScene) AddSceneEntities(entities ...shape.Shape) {
	s.shapes = append(s.shapes, entities...)
}

// It projects the world point on to the viewport and then to the canvas all at once
// cw , ch is canvas dimensions.
// vw , vh is viewport dimensions.
// d is the distance between the camera and viewport
// * Note: z is needs to be divide with the resultant Vec3 after multiplying with this mtx
func ProjectionViewport(d, cw, ch, vw, vh float32) homocoord.Mat3x4 {
	px := d * cw / vw
	py := d * ch / vh
	return homocoord.Mat3x4{
		px, 0, 0, 0, // x * d * cw / vw
		0, py, 0, 0, // y * d * ch / vh
		0, 0, 1, 0, // z (for w in perspective divide)
	}
}
