package main

import (
	"fmt"

	"github.com/unk1ndled/nier/src/unk"
)

func main() {

	val := -0.5
	fmt.Print(unk.Map(val, 0, 1, 0, 10))
}
