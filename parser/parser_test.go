package parser

import (
	"os"
	"testing"

	"github.com/charmbracelet/log"
)

func TestObjParser(t *testing.T) {
	parser := NewObjParser()
	file, _ := os.Open("/Users/ajaygupta/Desktop/computerGraphics/scene/models/cube.obj")

	parser.ChangeReader(file)
	model := parser.ReadModel()

	if model == nil {
		t.Error("Model not parsed")
	}

	log.Info("TestObjParser", model.Name(), model)
}

func TestParser(t *testing.T) {
	parser := NewSceneParser()
	file, _ := os.Open("/Users/ajaygupta/Desktop/computerGraphics/scene/rasterizaor.json")

	parser.ChangeReader(file)
	scene := parser.ReadScene()

	if scene == nil {
		t.Error("scene is nil")
	}
	if len(scene.Instances) == 0 || len(scene.Models) == 0 {
		t.Error("scene has less number of instance or models")
	}
	log.Info("parserd data", "data", scene)
}
