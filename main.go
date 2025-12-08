package main

import (
	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/editor"
)

func main() {
	log.Info("---ComputerGraphics---")

	editor.InitEditor()

	e := editor.NewEditor()
	defer e.DeInitEditor()

	e.HandleEvents()

}
