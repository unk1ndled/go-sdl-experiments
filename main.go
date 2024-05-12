package main

import (
	"fmt"

	clock "github.com/unk1ndled/nier/src/clock/logic"
)

func main() {
	clock := clock.NewClock()
	fmt.Print(clock)
}
