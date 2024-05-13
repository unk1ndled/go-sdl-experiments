package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/unk1ndled/nier/src/sdlutil"
	starfield "github.com/unk1ndled/nier/src/starfield/logic"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 500
	screenHeight = 500
	staramnt     = 500
)

var (
	pixels []byte
	black  = sdlutil.Color{0, 0, 0}
	strclr = sdlutil.Color{150, 150, 150}
)

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

	// look into this pixel format
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, screenWidth, screenHeight)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to Create TEXTURE : %s\n", err)
		os.Exit(3)
	}
	defer tex.Destroy()

	pixels = make([]byte, screenHeight*screenWidth*4)
	quit := false
	sc := sdlutil.NewSdlContext(pixels, screenWidth ,screenHeight)
	starfield := starfield.NewStarfield(staramnt, screenWidth, screenHeight)

	for !quit {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			if e.GetType() == sdl.QUIT {
				quit = true
			}
		}

		//reset screen
		for y := 0; y < screenHeight; y++ {
			for x := 0; x < screenWidth; x++ {
				sc.SetPixel(x, y, &black)
			}
		}

		starfield.Update(sc, &strclr)
		tex.Update(nil, unsafe.Pointer(&pixels[0]), 4*screenWidth)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(40)

	}
}
