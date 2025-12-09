package editor

import (
	"github.com/charmbracelet/log"
	"github.com/veandco/go-sdl2/sdl"
)

type EventHandler func(event sdl.Event)

func (e *Editor) HandleEvents(eventsToHandle map[uint32]EventHandler) {

	for e.State != Stopped {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			handler, ok := eventsToHandle[event.GetType()]
			if !ok {
				log.Debug("Event Occurred", event, "Not Handled")
				continue
			}

			log.Debug("Event Occurred", event, handler)
			handler(event)
		}

		sdl.Delay(33)
	}
}
