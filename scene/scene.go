package scene

import (
	"github.com/relaxgameing/computerGraphics/raytracing/geom"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/entity"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/light"
)

var (
	InitialCameraDirection geom.Vector = geom.Vector{WorldPoint: geom.WorldPoint{X: 0, Y: 0, Z: 1}}
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
			position:      geom.WorldPoint{X: 0, Y: 0, Z: 0},
			viewDirection: InitialCameraDirection,
			rotation:      geom.Rotation{Pitch: 0, Yaw: 0, Roll: 0},
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
	position geom.WorldPoint

	//* ViewDirection is a unit vector in the direction where the camera is seeing
	viewDirection geom.Vector

	//* Rotation is the current rotation of the camera
	rotation geom.Rotation
}

func (c *Camera) GetPosition() geom.WorldPoint {
	return c.position
}

func (c *Camera) GetDirection() geom.Vector {
	return c.viewDirection
}
func (c *Camera) SetDirection(newDirection geom.Vector) {
	c.viewDirection = *newDirection.UnitVector()
}
func (c *Camera) GetRotation() geom.Rotation {
	return c.rotation
}

func (c *Camera) RotateCameraBy(rotation geom.Rotation) {
	c.rotation.Pitch += rotation.Pitch
	c.rotation.Yaw += rotation.Yaw
	c.rotation.Roll += rotation.Roll

	c.recomputeViewVector()
}

func (c *Camera) recomputeViewVector() {
	c.SetDirection(*InitialCameraDirection.ToVector().Rotate(c.rotation).UnitVector())
}

func (c *Camera) MoveCameraBy(position geom.Vector) {
	c.position.X += position.X
	c.position.Y += position.Y
	c.position.Z += position.Z
}
