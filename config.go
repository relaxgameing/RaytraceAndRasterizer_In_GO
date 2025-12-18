package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/editor"
	"github.com/relaxgameing/computerGraphics/rasterization/geom"
	rsScene "github.com/relaxgameing/computerGraphics/rasterization/scene"
	"github.com/relaxgameing/computerGraphics/rasterization/scene/shape"
	rayScene "github.com/relaxgameing/computerGraphics/raytracing/scene"
)

const (
	RayTracing = iota
	Rasterization
)

type OptionRequirement struct {
	scene editor.Scene
}

func NewRayTracingRequirements() *OptionRequirement {
	scene := rayScene.NewScene("raytracing")
	return &OptionRequirement{
		scene: scene,
	}
}

func NewRasterizationRequirements() *OptionRequirement {
	scene := rsScene.NewScene()
	scene.AddSceneEntities(
		// shape.NewLine(geom.Point{-10, -100}, geom.Point{10, 100},
		// 	shape.WithColor(common.ColorBlue)),
		shape.NewLine(geom.Point{-200, -100}, geom.Point{240, 120},
			shape.WithColor(common.ColorBlue)),
		shape.NewLine(geom.Point{-50, -200}, geom.Point{60, 240},
			shape.WithColor(common.ColorRed)),
	)
	return &OptionRequirement{
		scene: scene,
	}
}

type RenderConfig struct {
	Technique int
	Options   *OptionRequirement
}

func handleArgs() RenderConfig {
	defaultConfig := RenderConfig{
		Technique: Rasterization,
		Options:   NewRasterizationRequirements(),
	}

	args := os.Args[1:]
	if len(args) == 0 {
		return defaultConfig
	}

	switch args[0] {
	case "ray":
		return RenderConfig{
			Technique: RayTracing,
			Options:   NewRayTracingRequirements(),
		}
	case "rs":
		return RenderConfig{
			Technique: Rasterization,
			Options:   NewRasterizationRequirements(),
		}
	default:
		log.Info("No proper Technique specified.\nEither ray for raytracing or rs for rasterization")
		return defaultConfig
	}
}
