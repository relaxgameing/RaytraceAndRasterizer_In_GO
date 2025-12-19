package shape

import (
	"math"

	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/rasterization/geom"
	"github.com/veandco/go-sdl2/sdl"
)

type LineOptions func(l *Line)

type Line struct {
	//* NOTE: start and end are just name there is no ordering
	start geom.Point
	end   geom.Point

	intercept float32
	slope     float32

	color sdl.Color
}

func NewLine(p1, p2 geom.Point, options ...LineOptions) *Line {
	slope := CalculateSlope(p1, p2)
	line := &Line{
		start:     p1,
		end:       p2,
		slope:     slope,
		intercept: CalculateIntercept(p1, slope),
		color:     common.ColorWhite,
	}
	for _, op := range options {
		op(line)
	}

	return line
}

// * Line options
func WithColor(color sdl.Color) LineOptions {
	return func(l *Line) {
		l.color = color
	}
}

// * Line helper functions
func CalculateSlope(p1, p2 geom.Point) float32 {
	return float32(p1.Y-p2.Y) / float32(p1.X-p2.X)
}

func CalculateIntercept(p1 geom.Point, slope float32) float32 {
	return float32(p1.Y) - slope*float32(p1.X)
}

// * Line functions
func (l *Line) GetSlope() float32 {
	return l.slope
}

func (l *Line) GetIntercept() float32 {
	return l.intercept
}

func (l *Line) GetPoints() []geom.Point {
	return []geom.Point{l.start, l.end}
}

func (l *Line) ComputeXForY(y int) float32 {
	return (float32(y) - l.intercept) / l.slope
}

func (l *Line) ComputeYForX(x int) float32 {
	return (float32(x)*l.slope + l.intercept)
}

// * shape Interface
func (l *Line) Draw() []*geom.Point {
	points := make([]*geom.Point, 0)
	points = append(points, &geom.Point{X: l.start.X, Y: l.start.Y})
	points = append(points, &geom.Point{X: l.end.X, Y: l.end.Y})
	// more vertical
	if math.Abs(float64(l.slope)) > 1 {
		l, r := geom.UpperPoint(l.start, l.end)
		temp := geom.InterpolateAlongLine(l.Y, l.X, r.Y, r.X)
		for _, p := range temp {
			points = append(points, &geom.Point{
				X: p.Y,
				Y: p.X,
			})
		}
		return points
	}

	left, right := geom.LeftPoint(l.start, l.end)
	temp := geom.InterpolateAlongLine(left.X, left.Y, right.X, right.Y)
	for _, p := range temp {
		points = append(points, &geom.Point{
			X: p.X,
			Y: p.Y,
		})
	}
	return points
}

func (l *Line) GetColor() sdl.Color {
	return l.color
}
