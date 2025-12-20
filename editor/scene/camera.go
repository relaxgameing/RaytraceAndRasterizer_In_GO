package scene

import "github.com/relaxgameing/computerGraphics/geom"

var (
	InitialCameraDirection geom.Vector = geom.Vector{WorldPoint: geom.WorldPoint{X: 0, Y: 0, Z: 1}}
)

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

func NewCamera(position geom.WorldPoint, viewDirection geom.Vector, rotation geom.Rotation) *Camera {
	return &Camera{
		position:      geom.WorldPoint{X: 0, Y: 0, Z: 0},
		viewDirection: InitialCameraDirection,
		rotation:      geom.Rotation{Pitch: 0, Yaw: 0, Roll: 0},
	}
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
