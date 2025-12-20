package scene

type Scene interface {
	SetSceneName(name string)
	GetCanvasWidth() int
	GetCanvasHeight() int
}
