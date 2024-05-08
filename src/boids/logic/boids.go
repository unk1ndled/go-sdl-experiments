package boids

import (
	"github.com/unk1ndled/nier/src/ds"
)

const (
	maxVelocity  = 20
	startvlocity = 10
	startacc     = 10
)

type Boid struct {
	position     *ds.Vector2D
	velocity     *ds.Vector2D
	acceleration *ds.Vector2D
}

func (b *Boid) GetPos() *ds.Vector2D {
	return b.position
}

func NewBoid(x, y int) *Boid {
	fx, fy := float64(x), float64(y)
	return &Boid{position: ds.NewVec2D(fx, fy), velocity: ds.RandomVec2D(startvlocity, startvlocity), acceleration: ds.RandomVec2D(startacc, startacc)}
}

func RandomBoid(x, y int) *Boid {
	return &Boid{position: ds.RandomVec2D(x, y), velocity: ds.RandomVec2D(startvlocity, startvlocity), acceleration: ds.RandomVec2D(startacc, startacc)}
}

func (b *Boid) Update() {
	b.position.Add(b.velocity)
	b.velocity.Add(b.acceleration)
	b.velocity.ClampMagnitude(float64(maxVelocity))
	b.acceleration.MultiplyByScalar(0)
	// fmt.Printf(" new pos %v boid %v", b.position, b)
}
