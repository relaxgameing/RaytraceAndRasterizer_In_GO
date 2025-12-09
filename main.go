package main

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/veandco/go-sdl2/sdl"
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
	})

}
