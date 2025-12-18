package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/editor"
	rsScene "github.com/relaxgameing/computerGraphics/rasterization/scene"
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
	return &OptionRequirement{
		scene: rayScene.NewScene("raytracing"),
	}
}

func NewRasterizationRequirements() *OptionRequirement {
	return &OptionRequirement{
		scene: rsScene.NewScene("rasterization"),
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
