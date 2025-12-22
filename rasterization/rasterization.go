package rasterization

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/editor"
	eScene "github.com/relaxgameing/computerGraphics/editor/scene"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
	"github.com/relaxgameing/computerGraphics/rasterization/scene"
	"github.com/veandco/go-sdl2/sdl"
)

func Rasterization(e *editor.Editor) {
	curScene := e.Scene.(*scene.RasterScene)

	setRendererDrawColor(e.Renderer, common.ColorWhite)
	e.Renderer.Clear()
	// for _, shape := range curScene.GetShapes() {
	// 	curColor := shape.GetColor()
	// 	for _, p := range shape.Draw() {
	// 		x, y := curScene.CanvasToSdl(p.X, p.Y)
	// 		intensityColor := common.ChangeColorIntensity(curColor, p.Intensity)
	// 		setRendererDrawColor(e.Renderer, intensityColor)
	// 		e.Renderer.DrawPoint(int32(x), int32(y))
	// 	}
	// }

	projectionMtx := scene.ProjectionViewport(float32(curScene.ViewPort.DistanceFromOrigin),
		float32(curScene.Canvas.Width),
		float32(curScene.Canvas.Height),
		float32(curScene.ViewPort.Width),
		float32(curScene.ViewPort.Height))

	for _, instance := range curScene.Instances {

		model := curScene.Models[instance.Name()]
		log.Info("Rasterization -> Drawing Model", "model", model.Name())

		for i := 0; i < model.TriangleCount(); i++ {
			triangle := model.TriangleAt(i)

			curColor := triangle.GetColor()
			setRendererDrawColor(e.Renderer, curColor)

			transformationMtx := modelTransformation(instance)

			transformedA := homocoord.Mat4MulVec4(transformationMtx, triangle.GetVertex(0))
			transformedB := homocoord.Mat4MulVec4(transformationMtx, triangle.GetVertex(1))
			transformedC := homocoord.Mat4MulVec4(transformationMtx, triangle.GetVertex(2))

			pa := homocoord.Mat3x4MulVec4(projectionMtx, transformedA)
			pb := homocoord.Mat3x4MulVec4(projectionMtx, transformedB)
			pc := homocoord.Mat3x4MulVec4(projectionMtx, transformedC)

			drawLine(e.Renderer, curScene, pa, pb)
			drawLine(e.Renderer, curScene, pa, pc)
			drawLine(e.Renderer, curScene, pc, pb)
		}
	}

	e.Renderer.Present()
}

func modelTransformation(model *eScene.ModelInstance) homocoord.Mat4 {
	scale := model.GetScale()
	translation := model.GetTranslation()
	rotation := model.GetRotation()

	scaleMtx := homocoord.Scale(scale.X, scale.Y, scale.Y)
	translationMtx := homocoord.Translation(translation.X, translation.Y, translation.Z)

	transformationMtx := homocoord.Mat4Mul(rotation, scaleMtx)
	return homocoord.Mat4Mul(translationMtx, transformationMtx)
	// return translationMtx
}

func drawLine(renderer *sdl.Renderer, curScene *scene.RasterScene, a, b homocoord.Vec3) {
	x1, y1 := curScene.CanvasToSdl(int(a.X/a.Z), int(a.Y/a.Z))
	x2, y2 := curScene.CanvasToSdl(int(b.X/b.Z), int(b.Y/b.Z))
	renderer.DrawLine(x1, y1, x2, y2)
}

func setRendererDrawColor(r *sdl.Renderer, color sdl.Color) {
	r.SetDrawColor(color.R, color.G, color.B, color.A)
}
