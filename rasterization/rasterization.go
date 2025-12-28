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
				drawTriangle(curScene, e.Renderer, tri, transformationMtx, projectionMtx)
			}

		}
	}

	curScene.ResetDepthBuffer()
	e.Renderer.Present()
}

func drawTriangle(curScene *scene.RasterScene, renderer *sdl.Renderer, tri geom.Triangle, transformationMtx homo.Mat4, projectionMtx homo.Mat3x4) {
	pa, pb, pc := transformAndProjectTriangle(tri, transformationMtx, projectionMtx)

	a := *geom.NewPoint(pa.X/pa.Z, pa.Y/pa.Z, pa.Z)
	b := *geom.NewPoint(pb.X/pb.Z, pb.Y/pb.Z, pb.Z)
	c := *geom.NewPoint(pc.X/pc.Z, pc.Y/pc.Z, pc.Z)

	//? back face culling
	if !isFacingTowardsCamera(a, b, c) {
		return
	}

	points := make([]*geom.Point, 0)
	ab := *geom.NewLine(a, b)
	bc := *geom.NewLine(b, c)
	ca := *geom.NewLine(c, a)
	points = append(points, ab.Draw()...)
	points = append(points, bc.Draw()...)
	points = append(points, ca.Draw()...)

	//? all the points are still in world coord
	points = append(points, tri.FillTriangle(
		a, b, c,
	)...)

	for _, pp := range points {
		if curScene.DepthBufferAt(int(pp.X), int(pp.Y)) < (1 / pp.Z) {
			setRendererDrawColor(renderer, tri.GetColor())
			drawPoint(renderer, curScene, pp.Vec3)
			curScene.SetDepthBufferAt(int(pp.X), int(pp.Y), 1/pp.Z)
		}
	}
}

// * Assumption: Clockwise faces
func isFacingTowardsCamera(a, b, c geom.Point) bool {
	ab := b.Subtract(a.Vec3)
	ac := c.Subtract(a.Vec3)

	normal := ab.Cross(ac).UnitVector()

	return a.UnitVector().Dot(normal) <= 0
}

func transformAndProjectTriangle(triangle geom.Triangle, transformationMtx homo.Mat4, projectionMtx homo.Mat3x4) (pa, pb, pc homo.Vec3) {

	pa = transformAndProjectPoint(triangle.GetVertex(0), transformationMtx, projectionMtx)
	pb = transformAndProjectPoint(triangle.GetVertex(1), transformationMtx, projectionMtx)
	pc = transformAndProjectPoint(triangle.GetVertex(2), transformationMtx, projectionMtx)

	return pa, pb, pc
}

// * returns point on canvas with z value from world coord
func transformAndProjectPoint(point geom.Point, transformationMtx homo.Mat4, projectionMtx homo.Mat3x4) homo.Vec3 {
	transformedPoint := homo.Mat4MulVec4(transformationMtx, homo.Vec3ToHomogeneous(point.Vec3))

	return homo.Mat3x4MulVec4(projectionMtx, transformedPoint)
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
