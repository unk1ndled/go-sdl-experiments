package ds

import (
	"fmt"
	"math"
	"math/rand"
)

type Vector2D [2]float64

func (vec Vector2D) String() string {
	return fmt.Sprintf("(%v, %v)", vec[0], vec[1])
}
func NewVec2D(x, y float64) *Vector2D {
	return &Vector2D{x, y}
}

func RandomVec2DPositive(maxx, maxy float64) *Vector2D {
	x := -rand.Float64() * (maxx)
	y := -rand.Float64() * (maxy)
	return NewVec2D(x, y)
}

func RandomVec2D(maxx, maxy float64) *Vector2D {
	x := rand.Float64() * (maxx)
	y := rand.Float64() * (maxy)

	if rand.Intn(2) == 1 {
		x *= -1
	}
	if rand.Intn(2) == 1 {
		y *= -1
	}

	return NewVec2D(x, y)
}

func (vec *Vector2D) Magnitude() float64 {
	return math.Sqrt(vec[0]*vec[0] + vec[1]*vec[1])
}

// get the director vector
func (vec *Vector2D) Normalized() *Vector2D {
	mag := vec.Magnitude()
	if mag == 0 {
		return &Vector2D{0, 0}
	}
	normalized := &Vector2D{vec[0] / mag, vec[1] / mag}
	return normalized
}

func (vec *Vector2D) SetMagnitude(value float64) {
	mag := vec.Magnitude()
	normalized := vec.Normalized()
	if mag == 0 {
		normalized = NewVec2D(1, 1)
	}

	vec[0] = normalized[0] * value
	vec[1] = normalized[1] * value

}

func (vec *Vector2D) ClampMagnitude(value float64) {
	if vec.Magnitude() > value {
		vec.SetMagnitude(value)
	}
}

func (vec *Vector2D) Add(other *Vector2D) {

	vec[0] += other[0]
	vec[1] += other[1]

}
func SubtractVectors(v1, v2 Vector2D) *Vector2D {
	return NewVec2D(v1[0]-v2[0], v1[1]-v2[1])
}
func (vec *Vector2D) Subtract(other *Vector2D) {

	vec[0] -= other[0]
	vec[1] -= other[1]

}

func (vec *Vector2D) Dist(other *Vector2D) float64 {
	dx := other[0] - vec[0]
	dy := other[1] - vec[1]
	return math.Sqrt(dx*dx + dy*dy)
}

func (vec *Vector2D) MultiplyByScalar(value float64) {
	vec[0] *= value
	vec[1] *= value
}
