package raytracing

import (
	"math"

	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/raytracing/scene"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/entity"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/light"
	"github.com/veandco/go-sdl2/sdl"
)

func RayTracing(e *editor.Editor) {
	curScene := e.Scene.(*scene.RayScene)

	for i := -curScene.Canvas.Width / 2; i <= curScene.Canvas.Width/2; i++ {
		for j := -curScene.Canvas.Height / 2; j <= curScene.Canvas.Height/2; j++ {
			curRay := generateViewPortRay(curScene, i, j)

			colorOfViewPort := traceRay(curRay, curScene, 3)

			e.Renderer.SetDrawColor(
				uint8(min(255, float32(colorOfViewPort.R))),
				uint8(min(255, float32(colorOfViewPort.G))),
				uint8(min(255, float32(colorOfViewPort.B))),
				uint8(min(255, float32(colorOfViewPort.A))))
			pi, pj := curScene.CanvasToSdl(i, j)
			e.Renderer.DrawPoint(int32(pi), int32(pj))
		}
	}
	e.Renderer.Present()
}

func ScalarProductColor(color sdl.Color, factor float32) *sdl.Color {
	return &sdl.Color{
		R: uint8(min(255, factor*float32(color.R))),
		G: uint8(min(255, factor*float32(color.G))),
		B: uint8(min(255, factor*float32(color.B))),
		A: uint8(min(255, factor*float32(color.A))),
	}
}

func generateViewPortRay(s *scene.RayScene, i int, j int) *geom.Ray {
	vx, vy := s.CanvasToViewPort(i, j)

	//todo: refactor the raytracer to use homocoord
	cameraPosition := s.ViewCamera.GetPosition()
	// cameraRotation := s.ViewCamera.GetRotation()

	var curRay geom.Ray = geom.Ray{
		Point:  geom.WorldPoint{X: 0, Y: 0, Z: 0},
		Lambda: 1e6,
		DirectionVector: *geom.NewVector(geom.WorldPoint{
			X: cameraPosition.X + vx,
			Y: cameraPosition.Y + vy,
			// Z: cameraPosition.Z + float32(s.ViewPort.DistanceFromOrigin)},
			// cameraPosition).Rotate(cameraRotation),
			Z: cameraPosition.Z},
			geom.WorldPoint{X: 0, Y: 0, Z: 0}),
	}

	return &curRay
}

func traceRay(ray *geom.Ray, curScene *scene.RayScene, rayDept int) (color sdl.Color) {
	var colorOfViewPort sdl.Color = sdl.Color{R: 0, G: 0, B: 0, A: 0}

	if rayDept == 0 {
		return sdl.Color{0, 0, 0, 0}
	}

	closestEntity, closestEntityLambda := getClosestEntityOnPathOfRay(ray, curScene.SceneEntities)
	if closestEntity == nil {
		return colorOfViewPort
	}

	targetPoint := ray.GetPointOnRayWithLambda(closestEntityLambda)
	normalVector := geom.NewVector(*targetPoint, closestEntity.GetOrigin())

	intensity := computeLightIntensityAtPoint(curScene, *targetPoint, *normalVector, closestEntity.GetSpecularExponent(), ray.Point)
	curColor := *ScalarProductColor(closestEntity.GetColor(), intensity)

	if closestEntity.GetReflectiveCoefficient() <= 0 {
		return curColor
	}

	reflectionRay := geom.Ray{
		Point:           *targetPoint,
		Lambda:          1,
		DirectionVector: *normalVector.MirrorReflectVector(*ray.DirectionVector.ScalarProduct(-1)),
	}
	reflectionColor := traceRay(&reflectionRay, curScene, rayDept-1)
	reflectiveness := closestEntity.GetReflectiveCoefficient()

	return sdl.Color{
		R: uint8(min(255, (1-reflectiveness)*float32(curColor.R)+
			reflectiveness*float32(reflectionColor.R))),
		G: uint8(min(255, (1-reflectiveness)*float32(curColor.G)+reflectiveness*float32(reflectionColor.G))),
		B: uint8(min(255, (1-reflectiveness)*float32(curColor.B)+reflectiveness*float32(reflectionColor.B))),
		A: uint8(min(255, (1-reflectiveness)*float32(curColor.A)+reflectiveness*float32(reflectionColor.A))),
	}

}

func getClosestEntityOnPathOfRay(
	ray *geom.Ray,
	sceneEntities []entity.Entity,
) (hitEntity entity.Entity, lambda float32) {

	var closestEntityLambda float32 = math.MaxFloat32
	var closestEntity entity.Entity
	for _, entity := range sceneEntities {
		if t, hit := entity.IsRayIntersecting(*ray); hit {
			if t < closestEntityLambda && t > light.Epsilon {
				closestEntityLambda = t
				closestEntity = entity
			}
		}
	}

	return closestEntity, closestEntityLambda
}

func computeLightIntensityAtPoint(
	scene *scene.RayScene,
	point geom.WorldPoint,
	normalVector geom.Vector,
	specular float32,
	cameraPosition geom.WorldPoint,
) float32 {

	var intensity float32 = 0
	for _, lighting := range scene.Lightings {

		if lighting.IsPointInFov(point, scene.SceneEntities) {
			intensity += lighting.ComputeDiffuseReflectionIntensityOfPoint(point, normalVector)
			intensity += lighting.ComputeSpecularReflectionIntensityOfPoint(point, normalVector, specular, cameraPosition)
		}
	}

	return intensity
}
