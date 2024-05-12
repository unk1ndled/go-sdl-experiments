package main

import (
	"fmt"
	"os"
	"unsafe"

	clock "github.com/unk1ndled/nier/src/clock/logic"
	"github.com/unk1ndled/nier/src/sdlutil"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	clockborder = 5

	screenWidth  = 500
	screenHeight = 300
	cellSize     = 10
	digitSize    = cellSize * 3
	digitSpace   = 5
)

var (
	pixels     []byte
	clockcolor = sdlutil.Color{0, 200, 0}
	black      = sdlutil.Color{0, 0, 0}
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to initialise SDL : %s\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("out of time", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
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
	clock := clock.NewClock()
	clock.Update()
	quit := false

	sc := sdlutil.NewSdlContext(pixels, screenWidth, screenHeight)

	for !quit {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			if e.GetType() == sdl.QUIT {
				quit = true
			}
		}
		for y := 0; y < screenHeight; y++ {
			for x := 0; x < screenWidth; x++ {
				if (x > clockborder && x < screenWidth-clockborder) && (y > clockborder && y < screenHeight-clockborder) {
					sc.SetPixel(x, y, &black)
				} else {
					sc.SetPixel(x, y, &clockcolor)
				}
			}
		}

		clock.Update()

		x := (screenWidth / 2) - (4 * digitSize)
		x += cellSize

		for i := 0; i < 6; i++ {
			sc.DrawDigit((*clock)[i], x, screenHeight/2, cellSize, &clockcolor)
			x += digitSize + digitSpace
			if i%2 == 1 && i < 5 {
				x -= (digitSize + digitSpace)
				dx := x + 2*digitSpace + 2*cellSize
				sc.DrawRect(dx-3, (screenHeight/2)-cellSize*3/2, cellSize, cellSize-1, &clockcolor)
				sc.DrawRect(dx-3, (screenHeight/2)+cellSize/2, cellSize, cellSize-1, &clockcolor)
				x += 2 * digitSize
			}
		}

		tex.Update(nil, unsafe.Pointer(&pixels[0]), 4*screenWidth)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(50)
	}

}
