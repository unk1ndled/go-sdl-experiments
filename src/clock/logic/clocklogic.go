package clock

import (
	"fmt"
	"time"
)

type Clock []byte

func (c Clock) String() string {
	return fmt.Sprintf(" %d%d : %d%d :%d%d ", c[0], c[1], c[2], c[3], c[4], c[5])
}

func NewClock() *Clock {
	clock := Clock{}
	for i := 0; i < 6; i++ {
		clock = append(clock, 0)
	}
	UpdateClock(clock)
	return &clock
}

func (c Clock) Update() {
	UpdateClock(c)
}
func UpdateClock(clock Clock) {
	currentTime := time.Now()
	hours := currentTime.Hour()
	minutes := currentTime.Minute()
	seconds := currentTime.Second()
	clock[0] = byte(hours / 10)
	clock[1] = byte(hours % 10)
	clock[2] = byte(minutes / 10)
	clock[3] = byte(minutes % 10)
	clock[4] = byte(seconds / 10)
	clock[5] = byte(seconds % 10)
}
