package parser

import (
	"os"
	"testing"

	"github.com/charmbracelet/log"
)

func TestObjParser(t *testing.T) {
	parser := NewObjParser()
	file, _ := os.Open("../models/cube.obj")

	parser.ChangeReader(file)
	model := parser.ReadModel()

	if model == nil {
		t.Error("Model not parsed")
	}

	log.Info("TestObjParser", model.name, model)
}

func TestParser(t *testing.T) {
	parser := NewSceneParser()
	file, _ := os.Open("../scene/rasterizaor.json")

	parser.ChangeReader(file)
	scene := parser.ReadScene()

	if scene == nil {
		t.Error("scene is nil")
	}
	if len(scene.instances) != 1 || len(scene.models) != 1 {
		t.Error("scene has less number of instance or models")
	}
	log.Info("parserd data", "data", scene)
}
