package main

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/raytracing"
	"github.com/relaxgameing/computerGraphics/scene/entity"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	ColorRed   = sdl.Color{R: 255, G: 0, B: 0, A: 255}
	ColorGreen = sdl.Color{R: 0, G: 255, B: 0, A: 255}
	ColorBlue  = sdl.Color{R: 0, G: 0, B: 255, A: 255}
)

func main() {
	log.Info("---ComputerGraphics---")

	editor.InitEditor()

	e := editor.NewEditor()
	defer e.DeInitEditor()

	e.Scene.AddEntity(
		entity.NewSphere(
			geom.WorldPoint{
				X: 0.0,
				Y: -1.0,
				Z: 3.0,
			},
			1,
			ColorRed,
		),
		entity.NewSphere(
			geom.WorldPoint{
				X: 2.0,
				Y: 0.0,
				Z: 4.0,
			},
			1,
			ColorBlue,
		),
		entity.NewSphere(
			geom.WorldPoint{
				X: -2.0,
				Y: 0.0,
				Z: 4.0,
			},
			1,
			ColorGreen,
		),
	)

	e.HandleEvents(map[uint32]editor.EventHandler{
		sdl.QUIT: func(event sdl.Event) {
			e.State = editor.Stopped
		},
		sdl.MOUSEBUTTONDOWN: func(event sdl.Event) {
			log.Info("RayTracing")
			raytracing.RayTracing(e)
		},
	})

}
