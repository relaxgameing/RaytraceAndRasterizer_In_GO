package scene

func (s *Scene) CanvasToViewPort(cx, cy int) (vx, vy float32) {
	vx = float32(cx) * float32(s.ViewPort.Width) / float32(s.Canvas.Width)
	vy = float32(cy) * float32(s.ViewPort.Height) / float32(s.Canvas.Height)

	return vx, vy
}

func (s *Scene) CanvasToSdl(cx, cy int) (x, y int) {
	x = (s.Canvas.Width / 2) + cx
	y = (s.Canvas.Height / 2) - cy

	return x, y
}
