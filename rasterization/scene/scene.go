package scene

type Scene struct {
	name string
}

func NewScene(name string) *Scene {
	return &Scene{
		name: name,
	}
}

func (s *Scene) SetSceneName(name string) {
	s.name = name
}
