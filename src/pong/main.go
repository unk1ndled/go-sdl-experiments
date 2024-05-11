package main

import (
	"fmt"
	"os"

	pong "github.com/unk1ndled/nier/src/pong/logic"
	"github.com/veandco/go-sdl2/sdl"
)

type Playable interface {
	Play()
}

type GameLoop struct {
	screenWidth  int32
	screenHeight int32
	game         Playable
}

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to initialise SDL : %s\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	gl := GameLoop{screenWidth: 1200, screenHeight: 400}
	window, err := sdl.CreateWindow("birdies", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, gl.screenWidth, gl.screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to Create window : %s\n", err)
		os.Exit(2)
	}
	defer window.Destroy()

	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	gl.game = pong.InitPong(gl.screenWidth, gl.screenHeight, renderer)
	gl.game.Play()
}
