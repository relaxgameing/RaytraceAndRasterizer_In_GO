package editor

func (e *Editor) canvasToViewPort(cx, cy int) (vx, vy int) {
	vx = cx * (e.ViewPort.Width / e.Canvas.Width)
	vy = cy * (e.ViewPort.Height / e.Canvas.Height)

	return vx, vy
}

func (e *Editor) canvasToSdl(cx, cy int) (x, y int) {
	x = (e.Canvas.Width / 2) + cx
	y = (e.Canvas.Height / 2) - cy

	return x, y
}
