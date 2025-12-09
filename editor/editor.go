package editor

import (
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/veandco/go-sdl2/sdl"
)

type EditorState int

const (
	Active EditorState = iota
	Error
	Stopped
)

/*
* Canvas:
* it is the screen which we are able to see in the compute
* it's unit is pixels
* it is a 2D canvas
 */

type Canvas struct {
	Width  int
	Height int
}

/*
*ViewPort:
* it is the window through which we see the real world
* it is world units
* it is a 3D world
 */

type ViewPort struct {
	Width              int
	Height             int
	DistanceFromOrigin int
}

type Editor struct {
	Id       uuid.UUID
	Window   *sdl.Window
	Renderer *sdl.Renderer

	Canvas
	ViewPort

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
		Canvas: Canvas{
			Width:  800,
			Height: 600,
		},
		ViewPort: ViewPort{
			Width:              1,
			Height:             1,
			DistanceFromOrigin: 1,
		},
	}
}
