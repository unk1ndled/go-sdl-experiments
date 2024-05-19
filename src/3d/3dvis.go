package main

import (
	"fmt"
	"math"
	"os"

	"github.com/unk1ndled/nier/src/3d/shapes"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 600
)

var (
	// pixels []byte
	strclr = sdl.Color{30, 30, 30, 255}
)

func smoothvalue(i, phase float64) float64 {
	return math.Sin(0.01*float64(i) + (phase * math.Pi / 3))
}

func randomRGB(i int) (uint8, uint8, uint8) {
	R := smoothvalue(float64(i), 0)*127 + 127
	G := smoothvalue(float64(i), 2)*127 + 127
	B := smoothvalue(float64(i), 4)*127 + 127
	return uint8(R), uint8(G), uint8(B)
}

func main() {

	//setup
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to initialise SDL : %s\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("stars", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to Create window : %s\n", err)
		os.Exit(2)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to Create Renderer : %s\n", err)
		os.Exit(3)
	}
	defer renderer.Destroy()

	cube := shapes.NewCube(6, screenWidth/2, screenHeight/2)
	colorfactor := 3

	quit := false
	for !quit {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			if e.GetType() == sdl.QUIT {
				quit = true
			}
		}
		colorfactor++

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		cube.Update()
		cube.Draw(renderer, strclr)
		strclr.R, strclr.G, strclr.B = randomRGB(colorfactor)
		strclr.R = 0

		renderer.Present()

		sdl.Delay(15)

	}
}
