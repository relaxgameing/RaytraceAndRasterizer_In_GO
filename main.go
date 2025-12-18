package main

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/relaxgameing/computerGraphics/rasterization"
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

	e.HandleEvents(map[uint32]editor.EventHandler{
		sdl.QUIT: func(event sdl.Event) {
			e.State = editor.Stopped
		},
		sdl.MOUSEBUTTONDOWN: func(event sdl.Event) {
			log.Info("Rasterization -- Starting")
			rasterization.Rasterization()
			log.Info("Rasterization -- Completed")
		},
	})

}
