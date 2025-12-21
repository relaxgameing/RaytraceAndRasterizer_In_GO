package parser

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/geom"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
)

type ObjParser struct {
	reader *bufio.Reader
}

func NewObjParser() *ObjParser {
	return &ObjParser{}
}

func (obj *ObjParser) ChangeReader(reader io.Reader) {
	obj.reader = bufio.NewReader(reader)
}

func (obj *ObjParser) ReadModel() *Model {
	model := EmptyModel()
	if obj.reader == nil {
		log.Error("ObjParser -> ReadModel", "err", errors.New("Reader is nil"))
		return nil
	}

	for {
		line, _, err := obj.reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Error("ObjParser -> ReadModel", "err", err)
			return nil
		}

		cmd := string(line)

		switch {
		case strings.HasPrefix(cmd, "v "):
			model.AddVertices(obj.parseVertex(cmd))
		case strings.HasPrefix(cmd, "f "):
			tri, err := (obj.parseFace(model, cmd))
			if err != nil {
				log.Error("ObjParser -> ReadModel -> parseFace", "err", err)
			}
			model.triangles = append(model.triangles, tri)
		}

	}

	return model
}

// v x y z [w]
func (obj *ObjParser) parseVertex(vertexLine string) homocoord.Vec4 {
	args := strings.Fields(vertexLine)
	x, _ := strconv.ParseFloat(args[1], 32)
	y, _ := strconv.ParseFloat(args[2], 32)
	z, _ := strconv.ParseFloat(args[3], 32)
	return homocoord.Vec4{
		X: float32(x), Y: float32(y), Z: float32(z), W: 1,
	}
}

// f 1 2 3 color				   # Triangle (3 verts) (right now only this works)
// f 1/1/1 2/2/2 3/3/3 color       # Triangle with UVs/normals
// f 1//1 2//2 3//3 color          # Triangle with normals only
// f 1/1 2/2 3/3 4/4 color         # Quad with UVs
func (obj *ObjParser) parseFace(model *Model, faceLine string) (geom.Triangle, error) {
	args := strings.Fields(faceLine)

	x, _ := strconv.Atoi(args[1])
	y, _ := strconv.Atoi(args[2])
	z, _ := strconv.Atoi(args[3])
	color := common.StringToSdlColor(string(args[3]))

	if len(model.vertices) < max(x, y, z) && min(x, y, z) < 1 {
		return geom.Triangle{}, errors.New("Vertex doesn't exists")
	}

	return *geom.NewTriangle(
		model.vertices[x],
		model.vertices[y],
		model.vertices[z],
		color,
	), nil

}
