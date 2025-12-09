package scene

// viewport_size = 1 x 1
// projection_plane_d = 1
// sphere {
//     center = (0, -1, 3)
//     radius = 1
//     color = (255, 0, 0)  # Red
// }
// sphere {
//     center = (2, 0, 4)
//     radius = 1
//     color = (0, 0, 255)  # Blue
// }
// sphere {
//     center = (-2, 0, 4)
//     radius = 1
//     color = (0, 255, 0)  # Green
// }

type Scene struct {
	Name string

	Canvas
	ViewPort

	Objects []SceneObject
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

type WorldPoint struct {
	x int
	y int
	z int
}

type SceneObject interface {
	GetOrigin() *WorldPoint
	IsPointIntersecting(point WorldPoint) bool
}
