package main

import (
	"fmt"

	"github.com/unk1ndled/nier/src/ds"
)

func main() {
	// vec := ds.NewVec2D(-1, 1)
	// fmt.Printf(" normalised vec %v \n ", vec.Normalized())
	// vec.MultiplyByScalar(-10)
	// fmt.Printf(" multiplied vec %v \n", vec)
	// vec.SetMagnitude(25)
	// fmt.Printf(" vec set magnitude %v \n", vec)
	// fmt.Printf(" get magnitude %f \n", vec.Magnitude())

	// vc1 := ds.NewVec2D(0, 0)
	vc2 := ds.NewVec2D(10, 10)
	vc3 := ds.NewVec2D(-10, 3)

	fmt.Printf(" diff %v \n", ds.SubtractVectors(*vc2, *vc3))

}
