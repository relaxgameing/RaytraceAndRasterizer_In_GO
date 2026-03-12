package scene

func (s *BaseScene) CanvasToViewPort(cx, cy int) (vx, vy float32) {
	vp := s.GetViewPort()
	canvas := s.GetCanvas()
	vx = float32(cx) * float32(vp.Width) / float32(canvas.Width)
	vy = float32(cy) * float32(vp.Height) / float32(canvas.Height)

	return vx, vy
}

func (s *BaseScene) CanvasToSdl(cx, cy int) (x, y int32) {
	x = int32((s.Canvas.Width / 2) + cx)
	y = int32((s.Canvas.Height / 2) - cy)

	return x, y
}
