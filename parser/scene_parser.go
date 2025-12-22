package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/editor/scene"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

type SceneParser struct {
	Parser
}

type SceneInstanceJson struct {
	Name        string         `json:"model"`
	Scale       [3]float32     `json:"scale"`
	Translation [3]float32     `json:"translation"`
	Rotation    homocoord.Mat4 `json:"rotation"`
}

type SceneJson struct {
	ModelPaths []string            `json:"modelPath"`
	Instances  []SceneInstanceJson `json:"instances"`
	Camera     struct {
		Position [3]float64 `json:"position"`
		Target   [3]float64 `json:"target"`
		Up       [3]float64 `json:"up"`
	} `json:"camera"`
}

func NewSceneParser() *SceneParser {
	return &SceneParser{}
}

func (s *SceneParser) ChangeReader(reader io.Reader) {
	s.reader = reader
}

func (s *SceneParser) ReadScene() *scene.SceneObjects {
	if s.reader == nil {
		log.Error("SceneParser -> ReadScene", "err", errors.New("Reader is nil"))
		return nil
	}

	sceneData := SceneJson{}
	decoder := json.NewDecoder(s.reader)
	err := decoder.Decode(&sceneData)

	if err != nil {
		log.Error("SceneParser -> ReadScene", "err", err)
		return nil
	}

	scene := scene.SceneObjects{Models: readModels(sceneData.ModelPaths)}

	scene.Instances = readModelInstances(scene.Models, sceneData.Instances)

	return &scene
}

func readModels(paths []string) map[string]*scene.Model {
	models := make(map[string]*scene.Model)
	objParser := NewObjParser()
	for _, path := range paths {
		f, err := os.Open(path)

		if err != nil {
			log.Error("SceneParser -> ReadScene", "err", fmt.Sprintf("Error opening file %s", path))
			return nil
		}
		defer f.Close()

		objParser.ChangeReader(f)
		model := objParser.ReadModel()

		if model == nil {
			log.Info("SceneParser -> ReadScene", "reading Model", "model is nil")
			return nil
		}

		models[model.Name()] = model
	}

	return models
}

func readModelInstances(models map[string]*scene.Model, instanceData []SceneInstanceJson) []*scene.ModelInstance {
	instances := make([]*scene.ModelInstance, 0)
	for _, instance := range instanceData {
		model := scene.NewModelInstance(
			instance.Name,
			models[instance.Name],
			homocoord.Vec3{X: instance.Scale[0], Y: instance.Scale[1], Z: instance.Scale[2]},
			homocoord.Vec3{X: instance.Translation[0], Y: instance.Translation[1], Z: instance.Translation[2]},
			instance.Rotation)

		instances = append(instances, &model)
	}

	return instances
}
