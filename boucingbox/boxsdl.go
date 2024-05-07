package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	startammount        = 1
	maxammount          = 1000
	screenHeight        = 600
	screenWidth         = 1000
	avrgboxsize         = 5
	sizevariationfactor = 3
	avrgspeed           = 10
)

var (
	boxes []*Box
)

type Box struct {
	x          int
	y          int
	color      sdl.Color
	xdirection bool
	ydirection bool
	speed      int
	width      int
	height     int
}

func (p Box) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.x, p.y)
}

func (b *Box) move() {
	// fmt.Printf(" moved boc %s \n", *b)
	movedirection(&b.xdirection, b.speed, &b.x)
	movedirection(&b.ydirection, b.speed, &b.y)
	boost := b.speed + rand.Intn(b.speed*2)
	if b.x+b.width >= screenWidth || b.x <= 0 {
		b.xdirection = !b.xdirection
		movedirection(&b.xdirection, boost, &b.x)
		b.handleCollision(true, !(b.x <= 0))
	}
	if b.y+b.height >= screenHeight || b.y <= 0 {
		b.ydirection = !b.ydirection
		movedirection(&b.ydirection, boost, &b.y)
		b.handleCollision(true, !(b.y <= 0))

	}
}

func (b *Box) handleCollision(ishorizontale, ispositive bool) {

	size := avrgboxsize + rand.Intn(sizevariationfactor*avrgboxsize)
	childx := b.x
	childy := b.y
	xdir := trueorfalse()
	ydir := trueorfalse()

	if ishorizontale {
		if ispositive {
			childx -= (size + b.width)
			xdir = false
		} else {
			childx += (size + b.width)
			xdir = true
		}
	} else {
		if ispositive {
			childy -= (size + b.height)
			ydir = true
		} else {
			childy += (size + b.height)
			ydir = false
		}
	}
	spawnbox(childx, childy, size, xdir, ydir)
}

func movedirection(xacc *bool, speed int, coordnate *int) {

	if *xacc {
		*coordnate += speed
	} else {
		*coordnate -= speed
	}
}

func spawnbox(x, y, size int, xdir, ydir bool) {
	if len(boxes) < maxammount {
		speed := 1 + int(avrgspeed*float64(avrgboxsize)/float64(size))
		color := newrandcolor()
		boxes = append(boxes, &Box{x: x, y: y, speed: speed, color: *color, xdirection: xdir, ydirection: ydir, width: size, height: size})

	}

}

func generateBoxes(amount int) {
	for i := 0; i < amount; i++ {
		x, y := rand.Intn(screenWidth), rand.Intn(screenWidth)
		size := avrgboxsize + rand.Intn(sizevariationfactor*avrgboxsize)

		spawnbox(x, y, size, trueorfalse(), trueorfalse())
	}
}

func newrandcolor() *sdl.Color {
	return &sdl.Color{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255))}
}
func trueorfalse() bool {
	return rand.Intn(2) == 1
}

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to initialise SDL : %s\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("moving box", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to Create window : %s\n", err)
		os.Exit(2)
	}
	defer window.Destroy()

	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	quit := false
	generateBoxes(startammount)

	for !quit {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			if e.GetType() == sdl.QUIT {
				quit = true
			}
			switch e := e.(type) {
			case *sdl.KeyboardEvent:
				fmt.Printf("%d", e.Keysym)
			}
		}

		// fmt.Printf("x is %d y is %d", box.x, box.y)

		// gives copy lmao
		// for _, box := range boxes {
		for i := range boxes {
			box := boxes[i]
			color := newrandcolor()
			renderer.SetDrawColor(color.R, color.G, color.B, 255)
			box.move()

			rect := sdl.Rect{X: int32(box.x), Y: int32(box.y), W: int32(box.width), H: int32(box.height)}
			renderer.FillRect(&rect)
		}

		renderer.Present()
		time.Sleep(50 * time.Millisecond)
	}

}
