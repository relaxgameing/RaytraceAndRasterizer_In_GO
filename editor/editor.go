package editor

import (
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/veandco/go-sdl2/sdl"
)

type Editor struct {
	Id       uuid.UUID
	Window   *sdl.Window
	Renderer *sdl.Renderer
}

func InitEditor() {
	sdl.Init(sdl.INIT_EVERYTHING)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func (e *Editor) DeInitEditor() {
	e.Renderer.Destroy()
	e.Window.Destroy()
}

func NewEditor() *Editor {
	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Error(err)
	}
	id, err := uuid.NewUUID()
	if err != nil {
		log.Error(err)
	}

	return &Editor{
		Id:       id,
		Window:   window,
		Renderer: renderer,
	}
}
