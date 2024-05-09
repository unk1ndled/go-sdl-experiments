package boids

import (
	"fmt"

	"github.com/unk1ndled/nier/src/ds"
)

const (
	maxVelocity = 5
	maxforce    = 0.8

	cohesionradius   = 2
	separationradius = 24
	alignmentradius  = 12
	startvlocity     = 5
	startacc         = 0
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
	return &Boid{position: ds.NewVec2D(fx, fy), velocity: ds.RandomVec2DPositive(startvlocity, startvlocity), acceleration: ds.RandomVec2DPositive(startacc, startacc)}
}

func RandomBoid(x, y int) *Boid {
	fx, fy := float64(x), float64(y)
	return &Boid{position: ds.RandomVec2D(fx, fy), velocity: ds.RandomVec2D(startvlocity, startvlocity), acceleration: ds.RandomVec2D(startacc, startacc)}
}

func (b *Boid) Flock(boids, snapshot []Boid) {
	alignment := ds.NewVec2D(0, 0)
	cohesion := ds.NewVec2D(0, 0)
	separation := ds.NewVec2D(0, 0)
	numa, nums, numc := 0, 0, 0

	for i := 0; i < len(snapshot); i++ {
		other := snapshot[i]
		dist := b.position.Dist(other.position)
		if b != &other && dist < alignmentradius {
			alignment.Add(other.velocity)
			numa++
		}
		if b != &other && dist < separationradius {

			//vector that points from the other to this boid
			diff := ds.SubtractVectors(*b.position, *other.position)
			//its influence is inversly proportional to this boid
			inverse := dist * dist
			if inverse == 0 {
				inverse = 1
			}
			diff.MultiplyByScalar(1 / inverse)
			// fmt.Printf("sepearation : %v ", separation)
			separation.Add(diff)
			nums++
		}
		if b != &other && dist < cohesionradius {
			cohesion.Add(other.position)
			numc++
		}
		// fmt.Printf("separation final : %v ", separation)
	}
	if numa > 0 {
		diva := float64(1 / numa)
		Process(alignment, b.velocity, diva, maxVelocity, maxforce)
	}
	if nums > 0 {
		divs := float64(1 / nums)
		Process(separation, b.velocity, divs, maxVelocity, maxforce)
	}
	if numc > 0 {
		divc := float64(1 / numc)
		Process(cohesion, b.position, divc, maxVelocity, maxforce)
		cohesion.Subtract(b.velocity)
	}
	b.acceleration.Add(alignment)
	// b.acceleration.Add(separation)
	// b.acceleration.Add(cohesion)

}

func Process(vec, sub *ds.Vector2D, scalar, mag, clamp float64) {
	vec.MultiplyByScalar(scalar)
	vec.SetMagnitude(mag)
	fmt.Printf(" mag is %f ", vec.Magnitude())
	vec.Subtract(sub)
	vec.ClampMagnitude(clamp)
}

func (b *Boid) Update() {
	b.position.Add(b.velocity)
	if b.velocity[0] > 0 || b.velocity[1] > 0 {
		print("error")
	}
	b.velocity.Add(b.acceleration)
	b.velocity.ClampMagnitude(float64(maxVelocity))
	b.acceleration.MultiplyByScalar(0)
}

func Copy(og []Boid) *[]Boid {
	copy := []Boid{}
	for _, v := range og {
		copy = append(copy, v)
	}
	return &copy
}
