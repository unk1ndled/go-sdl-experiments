package starfield

import (
	"math/rand"

	"github.com/unk1ndled/nier/src/sdlutil"
	"github.com/unk1ndled/nier/src/unk"
)

const (
	startSize         = 2
	maxSize   float64 = 5
)

var (
	screenHeight float64
	screenWidth  float64

	// centerpos *unk.Vector2D
)

type Star struct {
	position *unk.Vector2D
	z        float64
	pz       float64
}

func NewStar() *Star {
	r := rand.Float64() * (screenWidth)
	return &Star{
		position: unk.RandomVec2D(screenWidth/2, screenHeight/2),

		z:  r,
		pz: r,
	}
}

func (star *Star) Reset() {
	star.position.SetCoordinateToRandom(0, screenWidth/2)
	star.position.SetCoordinateToRandom(1, screenHeight/2)
	star.z = (screenWidth)
	star.pz = star.z

}

func (star *Star) Update(speed int) {
	star.z -= float64(speed)
	if star.z <= 0.0 {
		star.Reset()
	}
}

type Starfield struct {
	speed int
	stars []*Star
}

func NewStarfield(amnt, w, h int) *Starfield {
	stars := make([]*Star, amnt)
	screenWidth = float64(w)
	screenHeight = float64(h)

	// centerpos = unk.NewVec2D((screenWidth)/2, (screenHeight)/2)
	for i := 0; i < amnt; i++ {
		stars[i] = NewStar()
	}

	return &Starfield{stars: stars, speed: 10}
}

func (s *Starfield) AlterSpeed(isincrease bool) {
	if isincrease {
		s.speed += 2
	} else {
		s.speed -= 2
	}

}

func (s Starfield) Update(sc *sdlutil.SdlContext, strclr *sdlutil.Color) {
	for _, star := range s.stars {
		star.Update(s.speed)
		starx := star.position[0]
		stary := star.position[1]
		xdivisionResult := starx / star.z
		ydivisionResult := stary / star.z
		x := unk.Map(xdivisionResult, -1, 1, 0, screenWidth)
		y := unk.Map(ydivisionResult, -1, 1, 0, screenHeight)

		if x < 0 || y < 0 || x > screenWidth || y > screenHeight {
			continue
		}

		// lx := unk.Map(starx/star.pz, 0, 1, 0, screenWidth)
		// ly := unk.Map(stary/star.pz, 0, 1, 0, screenHeight)

		factor := 0.001 * ((screenWidth) - star.z)

		size := int(maxSize * factor)

		strclr.R += byte(x+y) % 255
		sc.DrawCircle(int(x), int(y), size, strclr)
		// sc.DrawLine(int(lx), int(ly), int(x), int(y), strclr)

		// sc.DrawRect(int(x), int(y), size, size, strclr)
	}
}
