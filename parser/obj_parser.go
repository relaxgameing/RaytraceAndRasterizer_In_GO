package parser

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/relaxgameing/computerGraphics/common"
	"github.com/relaxgameing/computerGraphics/editor/scene"
	"github.com/relaxgameing/computerGraphics/geom"
	homocoord "github.com/relaxgameing/computerGraphics/geom/homo_coord"
	"github.com/veandco/go-sdl2/sdl"
)

type ObjParser struct {
	reader   io.Reader
	curColor sdl.Color
}

func NewObjParser() *ObjParser {
	return &ObjParser{}
}

func (obj *ObjParser) ChangeReader(reader io.Reader) {
	obj.reader = reader
}

func (obj *ObjParser) ReadModel() *scene.Model {
	if obj.reader == nil {
		log.Error("ObjParser -> ReadModel", "err", errors.New("Reader is nil"))
		return nil
	}

	name := ""
	vertices := make([]homocoord.Vec3, 0)
	triangles := make([]geom.Triangle, 0)

	reader := bufio.NewReader(obj.reader)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Error("ObjParser -> ReadModel", "err", err)
			return nil
		}

		cmd := string(line)

		switch {
		case strings.HasPrefix(cmd, "usemtl"):
			obj.curColor = common.StringToSdlColor(strings.Fields(cmd)[1])
		case strings.HasPrefix(cmd, "name"):
			name = strings.Fields(cmd)[1]
		case strings.HasPrefix(cmd, "v "):
			vertices = append(vertices, (obj.parseVertex(cmd)))
		case strings.HasPrefix(cmd, "f "):
			tri, err := (obj.parseFace(vertices, cmd))
			if err != nil {
				log.Error("ObjParser -> ReadModel -> parseFace", "err", err)
			}
			triangles = append(triangles, tri)
		}

	}

	return scene.NewModel(name, vertices, triangles)
}

// v x y z [w]
func (obj *ObjParser) parseVertex(vertexLine string) homocoord.Vec3 {
	args := strings.Fields(vertexLine)
	x, _ := strconv.ParseFloat(args[1], 32)
	y, _ := strconv.ParseFloat(args[2], 32)
	z, _ := strconv.ParseFloat(args[3], 32)
	return homocoord.Vec3{
		X: float32(x), Y: float32(y), Z: float32(z),
	}
}

// f 1 2 3 				   # Triangle (3 verts) (right now only this works)
// f 1/1/1 2/2/2 3/3/3        # Triangle with UVs/normals
// f 1//1 2//2 3//3           # Triangle with normals only
// f 1/1 2/2 3/3 4/4          # Quad with UVs
func (obj *ObjParser) parseFace(vertices []homocoord.Vec3, faceLine string) (geom.Triangle, error) {
	args := strings.Fields(faceLine)

	x, _ := strconv.Atoi(args[1])
	y, _ := strconv.Atoi(args[2])
	z, _ := strconv.Atoi(args[3])

	if len(vertices) < max(x, y, z) && min(x, y, z) < 1 {
		return geom.Triangle{}, errors.New("Vertex doesn't exists")
	}

	return *geom.NewTriangle(
		*geom.NewPoint(vertices[x-1].X, vertices[x-1].Y, vertices[x-1].Z),
		*geom.NewPoint(vertices[y-1].X, vertices[y-1].Y, vertices[y-1].Z),
		*geom.NewPoint(vertices[z-1].X, vertices[z-1].Y, vertices[z-1].Z),
		obj.curColor,
	), nil

}
