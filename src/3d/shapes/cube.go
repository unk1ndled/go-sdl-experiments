package shapes

import (
	"math"

	"github.com/unk1ndled/nier/src/sdlutil"
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

type Cube struct {
	vertecies []*Point
	angle     float64
}

var edges = [12][2]int{
	{0, 1}, {1, 2}, {2, 3}, {3, 0},
	{4, 5}, {5, 6}, {6, 7}, {7, 4},
	{0, 4}, {1, 5}, {2, 6}, {3, 7},
}

func NewCube(x, y, z, a float64) *Cube {
	ver := []*Point{
		{-1, -1, -1},
		{1, -1, -1},
		{1, 1, -1},
		{-1, 1, -1},
		{-1, -1, 1},
		{1, -1, 1},
		{1, 1, 1},
		{-1, 1, 1},
	}
	return &Cube{vertecies: ver, angle: 0.012}
}

func toScreen(coor, zcoor, desiredmax float64) int32 {
	// return int32(desiredmax/2 + coor*factor*desiredmax)
	return int32(unk.Map(coor/(2+zcoor), -1, 1, 0, 800))
}

func (c *Cube) Update() {
	for i := 0; i < len(c.vertecies); i++ {
		ver := c.vertecies[i]
		// ver.rotateX(c.angle)
		ver.rotateY(c.angle)
		// ver.rotateZ(c.angle)

	}

}

func (c *Cube) Draw(renderer *sdl.Renderer, color *sdlutil.Color) {
	renderer.SetDrawColor(0, 155, 0, 255)

	for _, edge := range edges {
		vertex1 := c.vertecies[edge[0]]
		vertex2 := c.vertecies[edge[1]]
		x1, y1 := toScreen(vertex1.x, vertex1.z, float64(800)), toScreen(vertex1.y, vertex1.z, float64(800))
		x2, y2 := toScreen(vertex2.x, vertex2.z, float64(800)), toScreen(vertex2.y, vertex2.z, float64(800))

		renderer.DrawLine(x1, y1, x2, y2)

	}

}
