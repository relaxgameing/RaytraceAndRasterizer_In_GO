package rasterization

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/editor"
	eScene "github.com/relaxgameing/computerGraphics/editor/scene"
	viewfrustum "github.com/relaxgameing/computerGraphics/editor/scene/view_frustum"
	"github.com/relaxgameing/computerGraphics/geom"
	homo "github.com/relaxgameing/computerGraphics/geom/homo_coord"
	"github.com/relaxgameing/computerGraphics/rasterization/scene"
	"github.com/veandco/go-sdl2/sdl"
)

func Rasterization(e *editor.Editor) {
	curScene := e.Scene.(*scene.RasterScene)

	setRendererDrawColor(e.Renderer, common.ColorWhite)
	e.Renderer.Clear()

	projectionMtx := scene.ProjectionViewport(float32(curScene.ViewPort.DistanceFromOrigin),
		float32(curScene.Canvas.Width),
		float32(curScene.Canvas.Height),
		float32(curScene.ViewPort.Width),
		float32(curScene.ViewPort.Height))

	for _, instance := range curScene.Instances {
		boundingSphere := instance.GetBoundingSphere()
		condition := curScene.ViewFrustum.ObjectInsideFrustum(boundingSphere.GetOrigin(), boundingSphere.GetRadius())

		if condition == viewfrustum.OutsideVolume {
			continue
		}

		model := curScene.Models[instance.Name()]
		log.Info("Rasterization -> Drawing Model", "model", model.Name())
		transformationMtx := modelTransformation(instance, curScene.ViewCamera)

		for i := 0; i < model.TriangleCount(); i++ {
			triangle := model.TriangleAt(i)
			curColor := triangle.GetColor()
			setRendererDrawColor(e.Renderer, curColor)

			subTriangles := make([]geom.Triangle, 0)
			if condition == viewfrustum.PartiallyInside {
				subTriangles = append(subTriangles, curScene.ViewFrustum.TriangleInsideFrustum(triangle)...)
			} else {
				subTriangles = append(subTriangles, triangle)
			}

			for _, tri := range subTriangles {
				pa, pb, pc := transformAndProjectTriangle(tri, transformationMtx, projectionMtx)

				a := *geom.NewPointFromVec3(pa.ScalarPrd(1 / pa.Z))
				b := *geom.NewPointFromVec3(pb.ScalarPrd(1 / pb.Z))
				c := *geom.NewPointFromVec3(pc.ScalarPrd(1 / pc.Z))

				points := make([]*geom.Point, 0)
				ab := *geom.NewLine(a, b)
				bc := *geom.NewLine(b, c)
				ca := *geom.NewLine(c, a)
				points = append(points, ab.Draw()...)
				points = append(points, bc.Draw()...)
				points = append(points, ca.Draw()...)

				log.Info("Rasterization -> before fillTriangle", "len", len(points))
				points = append(points, tri.FillTriangle(
					a, b, c,
				)...)
				log.Info("Rasterization -> After fillTriangle", "len", len(points))
				for _, point := range points {
					drawPoint(e.Renderer, curScene, point.Vec3)
				}
			}

		}
	}

	e.Renderer.Present()
}

func transformAndProjectTriangle(triangle geom.Triangle, transformationMtx homo.Mat4, projectionMtx homo.Mat3x4) (pa, pb, pc homo.Vec3) {
	transformedA := homo.Mat4MulVec4(transformationMtx, homo.Vec3ToHomogeneous(triangle.GetVertex(0).Vec3))
	transformedB := homo.Mat4MulVec4(transformationMtx, homo.Vec3ToHomogeneous(triangle.GetVertex(1).Vec3))
	transformedC := homo.Mat4MulVec4(transformationMtx, homo.Vec3ToHomogeneous(triangle.GetVertex(2).Vec3))

	pa = homo.Mat3x4MulVec4(projectionMtx, transformedA)
	pb = homo.Mat3x4MulVec4(projectionMtx, transformedB)
	pc = homo.Mat3x4MulVec4(projectionMtx, transformedC)

	return pa, pb, pc
}

func modelTransformation(model *eScene.ModelInstance, camera *eScene.Camera) homo.Mat4 {
	modelTranslation := model.GetTranslation()
	translation := modelTranslation.Subtract(camera.GetPosition())

	modelRotation := model.GetRotation()
	rotation := homo.Mat4Mul(modelRotation, camera.GetRotation())

	scale := model.GetScale()
	scaleMtx := homo.Scale(scale.X, scale.Y, scale.Y)

	translationMtx := homo.Translation(translation.X, translation.Y, translation.Z)
	transformationMtx := homo.Mat4Mul(rotation, scaleMtx)
	return homo.Mat4Mul(translationMtx, transformationMtx)
}

func drawLine(renderer *sdl.Renderer, curScene *scene.RasterScene, a, b homo.Vec3) {
	x1, y1 := curScene.CanvasToSdl(int(a.X/a.Z), int(a.Y/a.Z))
	x2, y2 := curScene.CanvasToSdl(int(b.X/b.Z), int(b.Y/b.Z))
	renderer.DrawLine(x1, y1, x2, y2)
}

func drawPoint(renderer *sdl.Renderer, curScene *scene.RasterScene, a homo.Vec3) {
	x1, y1 := curScene.CanvasToSdl(int(a.X), int(a.Y))
	renderer.DrawPoint(x1, y1)
}

func setRendererDrawColor(r *sdl.Renderer, color sdl.Color) {
	r.SetDrawColor(color.R, color.G, color.B, color.A)
}
