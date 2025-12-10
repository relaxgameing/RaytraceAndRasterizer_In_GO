package raytracing

import (
	"math"

	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/scene"
	"github.com/relaxgameing/computerGraphics/scene/entity"
	"github.com/veandco/go-sdl2/sdl"
)

func RayTracing(e *editor.Editor) {
	curScene := e.Scene

	for i := -curScene.Canvas.Width / 2; i <= curScene.Canvas.Width/2; i++ {
		for j := -curScene.Canvas.Height / 2; j <= curScene.Canvas.Height/2; j++ {
			curRay := generateViewPortRay(curScene, i, j)

			closestEntity, closestEntityLambda := getClosestEntityOnPathOfRay(curRay, curScene.SceneEntities)

			var intensity float32 = 1
			var colorOfViewPort sdl.Color = sdl.Color{R: 255, G: 255, B: 255, A: 255}
			cameraPosition := geom.WorldPoint{X: 0, Y: 0, Z: 0}

			if closestEntity != nil {
				targetPoint := curRay.GetPointOnRayWithLambda(closestEntityLambda)
				normalVector := geom.NewVector(*targetPoint, closestEntity.GetOrigin())

				intensity = computeLightIntensityAtPoint(curScene, *targetPoint, *normalVector, closestEntity.GetSpecularExponent(), cameraPosition)
				colorOfViewPort = closestEntity.GetColor()
			}

			e.Renderer.SetDrawColor(
				uint8(min(255, intensity*float32(colorOfViewPort.R))),
				uint8(min(255, intensity*float32(colorOfViewPort.G))),
				uint8(min(255, intensity*float32(colorOfViewPort.B))),
				uint8(min(255, intensity*float32(colorOfViewPort.A))))
			pi, pj := curScene.CanvasToSdl(i, j)
			e.Renderer.DrawPoint(int32(pi), int32(pj))
		}
	}
	e.Renderer.Present()
}

func generateViewPortRay(s *scene.Scene, i int, j int) *geom.Ray {
	vx, vy := s.CanvasToViewPort(i, j)
	sceneOrigin := geom.WorldPoint{X: 0, Y: 0, Z: 0}

	var curRay geom.Ray = geom.Ray{
		Point:  sceneOrigin,
		Lambda: 1e6,
		DirectionVector: *geom.NewVector(geom.WorldPoint{
			X: vx,
			Y: vy,
			Z: float32(s.ViewPort.DistanceFromOrigin)},
			sceneOrigin),
	}

	return &curRay
}

func getClosestEntityOnPathOfRay(
	ray *geom.Ray,
	sceneEntities []entity.Entity,
) (hitEntity entity.Entity, lambda float32) {

	var closestEntityLambda float32 = math.MaxFloat32
	var closestEntity entity.Entity
	for _, entity := range sceneEntities {
		if t, hit := entity.IsRayIntersecting(*ray); hit {
			if t < closestEntityLambda {
				closestEntityLambda = t
				closestEntity = entity
			}
		}
	}

	return closestEntity, closestEntityLambda
}

func computeLightIntensityAtPoint(
	scene *scene.Scene,
	point geom.WorldPoint,
	normalVector geom.Vector,
	specular float32,
	cameraPosition geom.WorldPoint,
) float32 {

	var intensity float32 = 0
	for _, lighting := range scene.Lightings {
		intensity += lighting.ComputeDiffuseReflectionIntensityOfPoint(point, normalVector)
		intensity += lighting.ComputeSpecularReflectionIntensityOfPoint(point, normalVector, specular, cameraPosition)
	}

	return intensity
}
