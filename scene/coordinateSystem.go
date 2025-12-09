package scene

func (s *Scene) canvasToViewPort(cx, cy int) (vx, vy int) {
	vx = cx * (s.ViewPort.Width / s.Canvas.Width)
	vy = cy * (s.ViewPort.Height / s.Canvas.Height)

	return vx, vy
}

func (s *Scene) canvasToSdl(cx, cy int) (x, y int) {
	x = (s.Canvas.Width / 2) + cx
	y = (s.Canvas.Height / 2) - cy

	return x, y
}
