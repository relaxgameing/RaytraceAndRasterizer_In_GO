package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/editor/scene"
	"github.com/relaxgameing/computerGraphics/geom"
	"github.com/relaxgameing/computerGraphics/parser"
	rsScene "github.com/relaxgameing/computerGraphics/rasterization/scene"
	"github.com/relaxgameing/computerGraphics/rasterization/scene/shape"
	rayScene "github.com/relaxgameing/computerGraphics/raytracing/scene"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/entity"
	"github.com/relaxgameing/computerGraphics/raytracing/scene/light"
)

const (
	RayTracing = iota
	Rasterization
)

type OptionRequirement struct {
	scene scene.Scene
}

func NewRayTracingRequirements() *OptionRequirement {
	scene := rayScene.NewScene("raytracing")
	scene.AddLighting(
		light.NewAmbientLight(0.2, common.ColorWhite),
		light.NewDirectionalLight(*geom.WorldPoint{1, 4, 4}.ToVector(), 0.2, common.ColorWhite),
		light.NewPointLight(geom.WorldPoint{2, 1, 0}, 0.6, common.ColorWhite),
	)
	scene.AddSceneEntities(
		entity.NewSphere(
			geom.WorldPoint{0, -1, 3},
			1,
			common.ColorRed,
			500,
			0.2,
		),
		entity.NewSphere(
			geom.WorldPoint{0, -5001, 3},
			5000,
			common.ColorYellow,
			500,
			0.2,
		),
	)
	return &OptionRequirement{
		scene: scene,
	}
}

func NewRasterizationRequirements() *OptionRequirement {
	file, err := os.Open("./scene/rasterizaor.json")
	if err != nil {
		log.Error("NewRasterizationRequirements -> Error opening file", "err", err)
		return nil
	}
	sceneParser := parser.NewSceneParser()
	sceneParser.ChangeReader(file)

	scene := rsScene.NewRasterScene(
		rsScene.WithSceneObjects(sceneParser.ReadScene()),
	)
	scene.AddSceneEntities(
		shape.NewTriangle(
			*geom.NewPoint(-50, -200),
			*geom.NewPoint(50, -200),
			*geom.NewPoint(-50, 200),
		),
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
