package scene

import vf "github.com/relaxgameing/computerGraphics/editor/scene/view_frustum"

type Scene interface {
	GetName() string
	GetCanvas() *Canvas
	GetViewPort() *ViewPort
	GetCamera() *Camera
	GetViewFrustum() *vf.ViewFrustum
}

type BaseScene struct {
	Name string
	Canvas
	ViewPort

	ViewCamera  *Camera
	ViewFrustum vf.ViewFrustum
}

type SceneObjects struct {
	Models    map[string]*Model
	Instances []*ModelInstance
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

func (b *BaseScene) GetName() string                 { return b.Name }
func (b *BaseScene) GetCanvas() *Canvas              { return &b.Canvas }
func (b *BaseScene) GetViewPort() *ViewPort          { return &b.ViewPort }
func (b *BaseScene) GetCamera() *Camera              { return b.ViewCamera }
func (b *BaseScene) GetViewFrustum() *vf.ViewFrustum { return &b.ViewFrustum }
