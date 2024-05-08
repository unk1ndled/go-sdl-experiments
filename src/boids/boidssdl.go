package main

import (
	"fmt"
	"os"
	"time"

	"github.com/unk1ndled/nier/src/ds"

	boids "github.com/unk1ndled/nier/src/boids/logic"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 800
	boidamount   = 100
	boidsize     = 5
)

var (
	renderer *sdl.Renderer
)

type Sym struct {
	Boids     []boids.Boid
	Positions []*ds.Vector2D
	recs      []sdl.Rect
}

func NewSym() *Sym {
	sym := Sym{Boids: []boids.Boid{}, Positions: []*ds.Vector2D{}, recs: []sdl.Rect{}}
	for i := 0; i < boidamount; i++ {

		// int conversion issues
		boid := boids.NewBoid(int(screenWidth/2), int(screenHeight/2))
		sym.Boids = append(sym.Boids, *boid)
		sym.Positions = append(sym.Positions, boid.GetPos())
		x, y, w, h := int32(sym.Positions[i][0]), int32(sym.Positions[i][1]), int32(boidsize), int32(boidsize)
		sym.recs = append(sym.recs, sdl.Rect{X: x, Y: y, W: w, H: h})
	}
	return &sym
}

func (sym Sym) Update() {
	// print("hii")
	for _, boid := range sym.Boids {
		boid.Update()
	}

	for i := 0; i < len(sym.recs); i++ {
		// making boids appear inside the screen at all times
		sym.recs[i].X = (int32(sym.Positions[i][0]) + screenWidth) % screenWidth
		sym.recs[i].Y = (int32(sym.Positions[i][1]) + screenWidth) % screenHeight
	}
	sym.renderBoids()
}

func (sym Sym) renderBoids() {
	renderer.SetDrawColor(155, 0, 155, 255)
	renderer.FillRects(sym.recs)
}

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to initialise SDL : %s\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("cellulare automata", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Failed to Create window : %s\n", err)
		os.Exit(2)
	}
	defer window.Destroy()

	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	quit := false
	sym := NewSym()

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

		sym.Update()
		renderer.Present()
		time.Sleep(20 * time.Millisecond)
	}

}
