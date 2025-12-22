package common

import "github.com/veandco/go-sdl2/sdl"

var (
	ColorRed    = sdl.Color{R: 255, G: 0, B: 0, A: 255}
	ColorGreen  = sdl.Color{R: 0, G: 255, B: 0, A: 255}
	ColorBlue   = sdl.Color{R: 0, G: 0, B: 255, A: 255}
	ColorYellow = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	ColorCyan   = sdl.Color{R: 0, G: 255, B: 255, A: 255}
	ColorPurple = sdl.Color{R: 255, G: 0, B: 255, A: 255}
	ColorWhite  = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	ColorBlack  = sdl.Color{R: 0, G: 0, B: 0, A: 0}
)

func ChangeColorIntensity(color sdl.Color, intensity float32) sdl.Color {
	return sdl.Color{
		R: uint8(float32(color.R) * intensity),
		G: uint8(float32(color.G) * intensity),
		B: uint8(float32(color.B) * intensity),
		A: uint8(float32(color.A) * intensity),
	}
}

func StringToSdlColor(colorName string) sdl.Color {
	switch colorName {
	case "red":
		return ColorRed
	case "green":
		return ColorGreen
	case "blue":
		return ColorBlue
	case "yellow":
		return ColorYellow
	case "purple":
		return ColorPurple
	case "cyan":
		return ColorCyan
	case "white":
		return ColorWhite
	case "black":
		return ColorBlack
	default:
		return ColorBlack
	}
}
