package editor

type Scene interface {
	SetSceneName(name string)
	GetWidth() int
	GetHeight() int
}
