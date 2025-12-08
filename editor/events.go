package editor

import "github.com/veandco/go-sdl2/sdl"

func (e *Editor) HandleEvents() {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}

		sdl.Delay(33)
	}
}
