package main

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/relaxgameing/computerGraphics/rasterization"
	"github.com/relaxgameing/computerGraphics/raytracing"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	log.Info("---ComputerGraphics---")
	op := handleArgs()

	editor.InitEditor()

	e := editor.NewEditor(op.Options.scene)
	defer e.DeInitEditor()

	e.HandleEvents(map[uint32]editor.EventHandler{
		sdl.QUIT: func(event sdl.Event) {
			e.State = editor.Stopped
		},
		sdl.MOUSEBUTTONDOWN: func(event sdl.Event) {
			switch op.Technique {
			case RayTracing:
				log.Info("RayTracing -- Starting")
				raytracing.RayTracing(e)
				log.Info("RayTracing -- Completed")
			case Rasterization:
				log.Info("Rasterization -- Starting")
				rasterization.Rasterization(e)
				log.Info("Rasterization -- Completed")
			}
		},
		sdl.KEYDOWN: func(event sdl.Event) {
			keyEvent := event.(*sdl.KeyboardEvent)

			camera := e.Scene.GetCamera()
			switch sdl.GetKeyName(keyEvent.Keysym.Sym) {
			case "W":
				camera.MoveBy(camera.GetForwardDirection())
			case "S":
				dir := camera.GetForwardDirection()
				camera.MoveBy(dir.ScalarPrd(-1))
			case "A":
				forward := camera.GetForwardDirection()
				up := camera.GetUpDirection()
				leftDir := forward.Cross(up)
				camera.MoveBy(leftDir)
			case "D":
				forward := camera.GetForwardDirection()
				up := camera.GetUpDirection()
				right := forward.Cross(up)
				camera.MoveBy(right.ScalarPrd(-1))

			case "E":
				up := camera.GetUpDirection()
				camera.MoveBy(up)
			case "Q":
				up := camera.GetUpDirection()
				camera.MoveBy(up.ScalarPrd(-1))
			}

			rasterization.Rasterization(e)
		},
	})

}
