package raytracing

import (
	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/veandco/go-sdl2/sdl"
)

func RayTracing(e *editor.Editor) {
	scene := e.Scene
	sceneOrigin := geom.WorldPoint{X: 0, Y: 0, Z: 0}

	for i := -scene.Canvas.Width / 2; i <= scene.Canvas.Width/2; i++ {
		for j := -scene.Canvas.Height / 2; j <= scene.Canvas.Height/2; j++ {

			vx, vy := scene.CanvasToViewPort(i, j)

			var curRay geom.Ray = geom.Ray{
				Point:  sceneOrigin,
				Lambda: 1e6,
				DirectionVector: *geom.NewVector(geom.WorldPoint{
					X: vx,
					Y: vy,
					Z: float32(scene.ViewPort.DistanceFromOrigin)},
					sceneOrigin),
			}

			var colorOfViewPort sdl.Color = sdl.Color{R: 0, G: 0, B: 0, A: 255}
			for _, entity := range scene.SceneEntities {
				if entity.IsRayIntersecting(curRay) {
					colorOfViewPort = entity.GetColor()
				}
			}

			e.Renderer.SetDrawColor(colorOfViewPort.R, colorOfViewPort.G, colorOfViewPort.B, colorOfViewPort.A)
			pi, pj := scene.CanvasToSdl(i, j)
			e.Renderer.DrawPoint(int32(pi), int32(pj))
		}
	}
	e.Renderer.Present()
}
