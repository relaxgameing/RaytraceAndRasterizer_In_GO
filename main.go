package main

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/raytracing"
	"github.com/relaxgameing/computerGraphics/scene/entity"
	"github.com/relaxgameing/computerGraphics/scene/light"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	ColorRed    = sdl.Color{R: 255, G: 0, B: 0, A: 255}
	ColorGreen  = sdl.Color{R: 0, G: 255, B: 0, A: 255}
	ColorBlue   = sdl.Color{R: 0, G: 0, B: 255, A: 255}
	ColorYellow = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	ColorWhite  = sdl.Color{R: 255, G: 255, B: 255, A: 255}
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
			500,
			0.3,
		),
		entity.NewSphere(
			geom.WorldPoint{
				X: 2.0,
				Y: 0.0,
				Z: 4.0,
			},
			1,
			ColorBlue,
			500,
			0.3,
		),
		entity.NewSphere(
			geom.WorldPoint{
				X: -2.0,
				Y: 0.0,
				Z: 4.0,
			},
			1,
			ColorGreen,
			10,
			0.4,
		),
		entity.NewSphere(
			geom.WorldPoint{
				X: 0.0,
				Y: -5001.0,
				Z: 0.0,
			},
			5000,
			ColorYellow,
			1000,
			0.5,
		),
	)

	e.Scene.AddLighting(
		light.NewAmbientLight(0.2, ColorWhite),
		light.NewPointLight(geom.WorldPoint{X: 2, Y: 1, Z: 0}, 0.6, ColorWhite),
		light.NewDirectionalLight(geom.Vector{WorldPoint: geom.WorldPoint{X: 1, Y: 4, Z: 4}}, 0.2, ColorWhite),
	)

	e.HandleEvents(map[uint32]editor.EventHandler{
		sdl.QUIT: func(event sdl.Event) {
			e.State = editor.Stopped
		},
		sdl.MOUSEBUTTONDOWN: func(event sdl.Event) {
			log.Info("RayTracing - Starting")
			raytracing.RayTracing(e)
			log.Info("RayTracing - Complete")
		},
		sdl.KEYDOWN: func(event sdl.Event) {
			keyEvent := event.(*sdl.KeyboardEvent)

			log.Info("Keyboard Event Detected", "event", sdl.GetKeyName(keyEvent.Keysym.Sym))

			switch sdl.GetKeyName(keyEvent.Keysym.Sym) {
			case "W":
				e.Scene.MainCamera.Position.Z++
			case "A":
				e.Scene.MainCamera.Position.X--
			case "S":
				e.Scene.MainCamera.Position.Z--
			case "D":
				e.Scene.MainCamera.Position.X++
			case "Q":
				e.Scene.MainCamera.Position.Y--
			case "E":
				e.Scene.MainCamera.Position.Y++

			case "Left":
				e.Scene.MainCamera.Rotation.Yaw += geom.DegreeToRadian(-45)
			case "Right":
				e.Scene.MainCamera.Rotation.Yaw += geom.DegreeToRadian(45)
			case "Up":
				e.Scene.MainCamera.Rotation.Pitch += geom.DegreeToRadian(-45)
			case "Down":
				e.Scene.MainCamera.Rotation.Pitch += geom.DegreeToRadian(45)
			case "J":
				e.Scene.MainCamera.Rotation.Roll += geom.DegreeToRadian(45)
			case "L":
				e.Scene.MainCamera.Rotation.Roll += geom.DegreeToRadian(-45)
			}
			raytracing.RayTracing(e)
		},
	})

}
