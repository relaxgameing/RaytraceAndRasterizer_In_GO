package editor

import (
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/relaxgameing/computerGraphics/scene"
	"github.com/veandco/go-sdl2/sdl"
)

type EditorState int

const (
	Active EditorState = iota
	Error
	Stopped
)

type Editor struct {
	Id       uuid.UUID
	Window   *sdl.Window
	Renderer *sdl.Renderer

	Scene *scene.Scene

	State EditorState
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

// Todo: Dependency Injection for configuration of Editor
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
		State:    Active,
		Scene:    scene.NewScene("RayTracing"),
	}
}
