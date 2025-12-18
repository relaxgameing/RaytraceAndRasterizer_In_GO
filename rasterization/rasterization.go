package rasterization

import (
	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/relaxgameing/computerGraphics/rasterization/scene"
	"github.com/veandco/go-sdl2/sdl"
)

func Rasterization(e *editor.Editor) {
	curScene := e.Scene.(*scene.Scene)

	setRendererDrawColor(e.Renderer, common.ColorWhite)
	e.Renderer.Clear()
	for _, shape := range curScene.GetShapes() {
		curColor := shape.GetColor()
		setRendererDrawColor(e.Renderer, curColor)
		// x, y := curScene.CanvasToSdl(-50, -200)
		// a, b := curScene.CanvasToSdl(60, 240)
		// e.Renderer.DrawLine(int32(x), int32(y), int32(a), int32(b))
		for _, p := range shape.Draw() {
			x, y := curScene.CanvasToSdl(p.X, p.Y)
			e.Renderer.DrawPoint(int32(x), int32(y))
		}
	}

	e.Renderer.Present()
}

func setRendererDrawColor(r *sdl.Renderer, color sdl.Color) {
	r.SetDrawColor(color.R, color.G, color.B, color.A)
}
