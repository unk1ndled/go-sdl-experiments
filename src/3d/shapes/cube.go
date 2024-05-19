package shapes

import (
	"math"

	"github.com/unk1ndled/nier/src/unk"
	"github.com/veandco/go-sdl2/sdl"
)

type Point struct {
	x, y, z float64
}

func (pt *Point) rotateX(angle float64) {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	pt.y, pt.z = pt.y*cosA-pt.z*sinA, pt.y*sinA+pt.z*cosA
}

func (pt *Point) rotateY(angle float64) {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	pt.x, pt.z = pt.x*cosA+pt.z*sinA, -pt.x*sinA+pt.z*cosA
}

func (pt *Point) rotateZ(angle float64) {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	pt.x, pt.y = pt.x*cosA-pt.y*sinA, pt.x*sinA+pt.y*cosA
}

// cube logic
type Cube struct {
	coord unk.Vector2D
	scale float64

	xangle float64
	yangle float64
	zangle float64

	vertecies []*Point
}

// definining the edges for readability (which was the worst part abt ts)
var edges = [12][2]int{
	{0, 1}, {1, 2}, {2, 3}, {3, 0},
	{4, 5}, {5, 6}, {6, 7}, {7, 4},
	{0, 4}, {1, 5}, {2, 6}, {3, 7},
}

func NewCube(scale float64, x, y int) *Cube {
	ver := []*Point{
		{-1, -1, -1},
		{+1, -1, -1},
		{+1, +1, -1},
		{-1, +1, -1},
		{-1, -1, +1},
		{+1, -1, +1},
		{+1, +1, +1},
		{-1, +1, +1},
	}
	return &Cube{coord: *unk.NewVec2D(float64(x), float64(y)), vertecies: ver, xangle: -0.012, yangle: 0.012, zangle: -0.04, scale: scale}
}

func toScreen(coor, zcoor, scale, desiredmax float64) int32 {
	return int32(unk.Map(coor/(zcoor+scale), -1, 1, 0, desiredmax))
}

func (c *Cube) Update() {
	for i := 0; i < len(c.vertecies); i++ {
		ver := c.vertecies[i]
		ver.rotateX(c.xangle)
		ver.rotateY(c.yangle)
		ver.rotateZ(c.zangle)
	}

}

func (c *Cube) Draw(renderer *sdl.Renderer, color sdl.Color) {
	renderer.SetDrawColor(color.R, color.G, color.B, 255)
	scale := c.scale
	coord := c.coord
	for _, edge := range edges {
		vertex1 := c.vertecies[edge[0]]
		vertex2 := c.vertecies[edge[1]]
		x1, y1 := toScreen(vertex1.x, vertex1.z, scale, 2*coord[0]), toScreen(vertex1.y, vertex1.z, scale, 2*coord[1])
		x2, y2 := toScreen(vertex2.x, vertex2.z, scale, 2*coord[0]), toScreen(vertex2.y, vertex2.z, scale, 2*coord[1])
		renderer.DrawLine(x1, y1, x2, y2)
	}
}
