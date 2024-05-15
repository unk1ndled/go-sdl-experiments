package shapes

import (
	"github.com/unk1ndled/nier/src/sdlutil"
	"github.com/unk1ndled/nier/src/unk"
)

type Point struct {
	x, y, z float64
}

type Cube struct {
	vertecies []*Point
}

func NewCube(x, y, z, a float64) *Cube {
	ver := []*Point{
		{1, -1, -5},
		{1, -1, -3},
		{1, 1, -5},
		{1, 1, -3},
		{-1, -1, -5},
		{-1, -1, -3},
		{-1, 1, -5},
		{-1, 1, -3},
	}
	return &Cube{vertecies: ver}
}

func (c *Cube) Draw(sc *sdlutil.SdlContext, color *sdlutil.Color) {
	// for i := 0; i < 4; i++ {
	// 	right := c.vertecies[i]
	// 	left := c.vertecies[i+4]
	// 	x1 := unk.Map(right.x/right.z, -1, 1, 0, float64(sc.ScreenWidth))
	// 	x2 := unk.Map(left.x/left.z, -1, 1, 0, float64(sc.ScreenWidth))
	// 	y1 := unk.Map(right.y/right.z, -1, 1, 0, float64(sc.ScreenHeight))
	// 	y2 := unk.Map(left.y/left.z, -1, 1, 0, float64(sc.ScreenHeight))
	// 	sc.DrawLine(int(x1), int(y1), int(x2), int(y2), color)
	// }

	back := c.vertecies[0]
	front := c.vertecies[1]
	x1 := unk.Map(back.x/back.z, -1, 1, 0, float64(sc.ScreenWidth))
	x2 := unk.Map(front.x/front.z, -1, 1, 0, float64(sc.ScreenWidth))
	y1 := unk.Map(back.y/back.z, -1, 1, 0, float64(sc.ScreenHeight))
	y2 := unk.Map(front.y/front.z, -1, 1, 0, float64(sc.ScreenHeight))
	sc.DrawLine(int(x1), int(y1), int(x2), int(y2), color)

}
