package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/unk1ndled/nier/src/3d/shapes"
	"github.com/unk1ndled/nier/src/sdlutil"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 800
	staramnt     = 1000
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
	sc := sdlutil.NewSdlContext(pixels, screenWidth, screenHeight)

	cube := shapes.NewCube(screenWidth/2, screenHeight/2, 50, screenWidth)

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
		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_UP] != 0 {
			// starfield.AlterSpeed(true)
		} else if keys[sdl.SCANCODE_DOWN] != 0 {
			// starfield.AlterSpeed(false)
		}

		cube.Draw(sc, &strclr)

		tex.Update(nil, unsafe.Pointer(&pixels[0]), 4*screenWidth)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		sdl.Delay(100)

	}
}
