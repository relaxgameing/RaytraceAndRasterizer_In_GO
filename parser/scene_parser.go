package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/log"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

type SceneParser struct {
	Parser
}

type Scene struct {
	models    map[string]*Model
	instances []*ModelInstance
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

func (s *SceneParser) ReadScene() *Scene {
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

	scene := Scene{models: readModels(sceneData.ModelPaths)}

	scene.instances = readModelInstances(scene.models, sceneData.Instances)

	return &scene
}

func readModels(paths []string) map[string]*Model {
	models := make(map[string]*Model)
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

		models[model.name] = model
	}

	return models
}

func readModelInstances(models map[string]*Model, instanceData []SceneInstanceJson) []*ModelInstance {
	instances := make([]*ModelInstance, 0)
	for _, instance := range instanceData {
		modelInstance := ModelInstance{
			model:       models[instance.Name],
			name:        instance.Name,
			scale:       homocoord.Vec3{X: instance.Scale[0], Y: instance.Scale[1], Z: instance.Scale[2]},
			translation: homocoord.Vec3{X: instance.Translation[0], Y: instance.Translation[1], Z: instance.Translation[2]},
			rotation:    instance.Rotation,
		}

		instances = append(instances, &modelInstance)
	}

	return instances
}
