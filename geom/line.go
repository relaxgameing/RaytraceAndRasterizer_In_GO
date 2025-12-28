package geom

import (
	"math"

	"github.com/relaxgameing/computerGraphics/common"
	"github.com/veandco/go-sdl2/sdl"
)

type LineOptions func(l *Line)

type Line struct {
	//* NOTE: start and end are just name there is no ordering
	start Point
	end   Point

	intercept float32
	slope     float32

	color sdl.Color
}

func NewLine(p1, p2 Point, options ...LineOptions) *Line {
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
func CalculateSlope(p1, p2 Point) float32 {
	return float32(p1.Y-p2.Y) / float32(p1.X-p2.X)
}

func CalculateIntercept(p1 Point, slope float32) float32 {
	return float32(p1.Y) - slope*float32(p1.X)
}

// * Line functions
func (l *Line) GetSlope() float32 {
	return l.slope
}

func (l *Line) GetIntercept() float32 {
	return l.intercept
}

func (l *Line) GetPoints() []Point {
	return []Point{l.start, l.end}
}

func (l *Line) ComputeXForY(y int) float32 {
	return (float32(y) - l.intercept) / l.slope
}

func (l *Line) ComputeYForX(x int) float32 {
	return (float32(x)*l.slope + l.intercept)
}

// * shape Interface
func (l *Line) Draw() []*Point {
	points := make([]*Point, 0)
	points = append(points, NewPoint(l.start.X, l.start.Y, l.start.Z))
	points = append(points, NewPoint(l.end.X, l.end.Y, l.end.Z))
	// more vertical
	if math.Abs(float64(l.start.X-l.end.X)) < math.Abs(float64(l.start.Y-l.end.Y)) {
		upper, lower := UpperPoint(l.start, l.end)
		lineXVal := InterpolateAlongLine(lower.Y, lower.X, upper.Y, upper.X)
		lineZVal := InterpolateAlongLine(lower.Y, lower.Z, upper.Y, upper.Z)
		intensities := InterpolateAlongLine(float32(lower.Y), lower.Intensity, float32(upper.Y), upper.Intensity)
		for i, p := range lineXVal {
			points = append(points, NewPoint(
				p,
				lower.Y+float32(i),
				lineZVal[i],
				PointWithIntensity(intensities[i]),
			))
		}
		return points
	}

	left, right := LeftPoint(l.start, l.end)
	for i := 0; i < int(right.X-left.X-1); i++ {
		t := float32(i) / (right.X - left.X - 1)
		// p := LerpOnLine(left.Vec3, right.Vec3, t)
		lineYVal := InterpolateAlongLine(left.X, left.Y, right.X, right.Y)
		lineZVal := InterpolateAlongLine(left.X, left.Z, right.X, right.Z)
		intensity := Lerp(left.Intensity, right.Intensity, t)
		points = append(points, NewPoint(
			left.X+float32(i),
			lineYVal[i],
			lineZVal[i],
			PointWithIntensity(intensity)))
	}
	return points
}

func (l *Line) GetColor() sdl.Color {
	return l.color
}
