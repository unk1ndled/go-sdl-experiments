package boids

import (
	"github.com/unk1ndled/nier/src/ds"
)

const (
	maxVelocity = 5
	maxforce    = 0.3
	turnfactor  = 0.12

	separationWeight = 0.42
	alignmentWeight  = 0.20
	cohesionWeight   = 0.15

	// separationWeight = 0
	// alignmentWeight  = 0
	// cohesionWeight   = 0

	protectedRange = 15
	grouping       = 46
	alignmentrange = 50

	startvlocity = 2
	startacc     = 0
)

var (
	top    float64 = 20
	bottom float64 = 20
	left   float64 = 20
	right  float64 = 20
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
	fx, fy := float64(x), float64(y)
	return &Boid{position: ds.RandomVec2D(fx, fy), velocity: ds.RandomVec2D(startvlocity, startvlocity), acceleration: ds.RandomVec2D(startacc, startacc)}
}

func (b *Boid) AvoidWalls(w, h int32) {

	steer := ds.NewVec2D(0, 0)
	if b.position[0] < left {
		steer[0] += turnfactor

	} else if b.position[0] > float64(w)-right {
		steer[0] -= turnfactor

	}
	if b.position[1] < top {
		steer[1] += turnfactor

	} else if b.position[1] > float64(h)-bottom {
		steer[1] -= turnfactor
	}

	b.acceleration.Add(steer)

}

func (b *Boid) Alignment(snapshot []Boid) *ds.Vector2D {
	avgVelocity := ds.NewVec2D(0, 0)
	count := 0

	for _, otherBoid := range snapshot {
		if otherBoid != *b {
			distance := b.position.Dist(otherBoid.position)
			if distance < alignmentrange {
				avgVelocity.Add(otherBoid.velocity)
				count++
			}
		}
	}

	if count > 0 {
		avgVelocity.MultiplyByScalar(1 / float64(count))
		avgVelocity.SetMagnitude(maxVelocity)
		avgVelocity.Subtract(b.velocity)
		avgVelocity.ClampMagnitude(maxforce)
	}

	return avgVelocity
}

func (boid *Boid) Cohesion(snapshot []Boid) *ds.Vector2D {
	centerOfMass := ds.NewVec2D(0, 0)
	count := 0

	for _, otherBoid := range snapshot {
		if otherBoid != *boid {
			distance := boid.position.Dist(otherBoid.position)
			if distance < grouping {
				centerOfMass.Add(otherBoid.position)
				count++
			}
		}
	}

	if count > 0 {
		centerOfMass.MultiplyByScalar(1 / float64(count))
		desired := ds.SubtractVectors(*centerOfMass, *boid.position)
		desired.SetMagnitude(maxVelocity)
		steer := ds.SubtractVectors(*desired, *boid.velocity)
		steer.ClampMagnitude(maxforce)
		return steer
	}

	return ds.NewVec2D(0, 0)
}

func (b *Boid) Separation(snapshot []Boid) *ds.Vector2D {
	steer := ds.NewVec2D(0, 0)
	count := 0

	for _, otherBoid := range snapshot {
		if otherBoid != *b {
			distance := b.position.Dist(otherBoid.position)
			if distance < protectedRange {
				diff := ds.SubtractVectors(*b.position, *otherBoid.position)
				diff.SetMagnitude(1 / distance * distance)
				steer.Add(diff)
				count++
			}
		}
	}

	if count > 0 {

		// steer.MultiplyByScalar(float64(1 / count))
		steer.SetMagnitude(maxVelocity)
		steer.Subtract(b.velocity)
		steer.ClampMagnitude(maxforce)
	}

	return steer

}

func (b *Boid) Flock(snapshot []Boid) {
	ali := b.Alignment(snapshot)
	coh := b.Cohesion(snapshot)
	sep := b.Separation(snapshot)

	sep.MultiplyByScalar(separationWeight)
	ali.MultiplyByScalar(alignmentWeight)
	coh.MultiplyByScalar(cohesionWeight)

	b.acceleration.Add(sep)
	b.acceleration.Add(ali)
	b.acceleration.Add(coh)

}

func (b *Boid) Update(width, height int32) {
	b.position.Add(b.velocity)

	// stuck inside the screen
	// if b.position[0] < 0 || b.position[0] >= float64(width) {
	// 	b.position[0] = float64(int32(b.position[0]+float64(width)) % width)
	// }

	// if b.position[1] < 0 || b.position[1] >= float64(height) {
	// 	b.position[1] = float64(int32(b.position[1]+float64(height)) % height)
	// }

	b.AvoidWalls(width, height)
	b.velocity.Add(b.acceleration)

	b.velocity.ClampMagnitude(float64(maxVelocity))
	b.acceleration.MultiplyByScalar(0)
}

func Copy(og []Boid) []Boid {
	copy := []Boid{}
	copy = append(copy, og...)
	return copy
}
