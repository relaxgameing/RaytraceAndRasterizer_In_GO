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
	})

}
