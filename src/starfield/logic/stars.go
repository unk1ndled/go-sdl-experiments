package starfield

import (
	"math/rand"

	"github.com/unk1ndled/nier/src/sdlutil"
	"github.com/unk1ndled/nier/src/unk"
)

const (
	startSize         = 2
	maxSize   float64 = 7
	speed             = 2
)

var (
	screenHeight float64
	screenWidth  float64

	centerpos *unk.Vector2D
)

type Star struct {
	position *unk.Vector2D
	z        float64
}

func NewStar() *Star {
	return &Star{
		position: unk.RandomVec2D(screenWidth/2, screenHeight/2),

		z: rand.Float64() * (screenWidth),
	}
}

func (star *Star) Reset() {
	star.position.SetCoordinateToRandom(0, screenWidth/2)
	star.position.SetCoordinateToRandom(1, screenHeight/2)
	star.z = (screenWidth)

}

func (star *Star) Update() {
	star.z -= speed
	if star.z <= 0.0 {
		star.Reset()
	}
}

type Starfield struct {
	stars []*Star
}

func NewStarfield(amnt, w, h int) *Starfield {
	stars := make([]*Star, amnt)
	screenWidth = float64(w)
	screenHeight = float64(h)

	centerpos = unk.NewVec2D((screenWidth)/2, (screenHeight)/2)
	for i := 0; i < amnt; i++ {
		stars[i] = NewStar()
	}

	return &Starfield{stars: stars}
}

func (s Starfield) Update(sc *sdlutil.SdlContext, strclr *sdlutil.Color) {
	for _, star := range s.stars {
		star.Update()
		starx := star.position[0]
		stary := star.position[1]


		x := centerpos[0] + unk.Map(starx/star.z, 0, 1, 0, screenWidth/2)
		y := centerpos[1] + unk.Map(stary/star.z, 0, 1, 0, screenHeight/2)

		factor := 0.001 * ((screenWidth) - star.z)

		size := int(maxSize * factor)

		sc.DrawRect(int(x), int(y), size, size, strclr)
	}
}
