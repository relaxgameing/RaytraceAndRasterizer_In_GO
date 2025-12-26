package scene

import (
	"github.com/relaxgameing/computerGraphics/geom"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

var (
	InitialCameraDirection geom.Vector = geom.Vector{WorldPoint: geom.WorldPoint{X: 0, Y: 0, Z: 1}}
)

/*
*	Camera:
*	it is our eye through which we can see the scene
 */
type Camera struct {
	//* this is the position of the camera in the scene
	position homocoord.Vec3

	//* target represent the direction which the camera is looking
	target homocoord.Vec3
	//* this is the up direction from target
	up homocoord.Vec3

	//* Rotation is the current rotation of the camera
	rotation homocoord.Mat4
}

func NewCamera(position, target, up homocoord.Vec3, rotation homocoord.Mat4) *Camera {
	return &Camera{
		position: position,
		target:   target,
		up:       up,
		rotation: rotation,
	}
}

func (c *Camera) GetPosition() homocoord.Vec3 {
	return c.position
}

func (c *Camera) GetForwardDirection() homocoord.Vec3 {
	return c.target
}

func (c *Camera) GetUpDirection() homocoord.Vec3 {
	return c.up
}

func (c *Camera) GetRightDirection() homocoord.Vec3 {
	return c.target.Cross(c.up)
}

func (c *Camera) GetRotation() homocoord.Mat4 {
	return c.rotation
}

func (c *Camera) MoveBy(translation homocoord.Vec3) {
	c.position = c.position.Add(translation)
}

func (c *Camera) RotateBy(rotation homocoord.Mat4) {
	c.rotation = homocoord.Mat4Mul(c.rotation, rotation)
}
