package main

import (
	"fmt"
	"os"
	"time"

	gameoflife "github.com/unk1ndled/nier/src/gameoflife/logic"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 400
	screenHeight = 400
	cellSize     = 5
)

type Game struct {
	Renderer *sdl.Renderer
	Board    *gameoflife.Board
}

func NewGame(renderer *sdl.Renderer, board *gameoflife.Board) *Game {
	return &Game{
		Renderer: renderer,
		Board:    board,
	}
}

func (g *Game) DrawGrid() {
	g.Renderer.SetDrawColor(0, 0, 0, 255)
	for i := 0; i < screenWidth; i += cellSize {
		g.Renderer.DrawLine(int32(i), 0, int32(i), int32(screenHeight))
		g.Renderer.DrawLine(0, int32(i), int32(screenWidth), int32(i))
	}
}

func (g *Game) DrawRectangle(x, y int32) {
	rect := sdl.Rect{X: x, Y: y, W: cellSize, H: cellSize}
	g.Renderer.SetDrawColor(60, 6, 122, 100)
	g.Renderer.FillRect(&rect)
}

func (g *Game) DrawLiveCells() {
	g.Renderer.SetDrawColor(0, 0, 0, 255)
	for _, pair := range g.Board.Livecells {
		g.DrawRectangle(int32(pair[0]*cellSize), int32(pair[1]*cellSize))
	}
}

func (g *Game) Run() {
	running := true
	lastFrameTime := time.Now()
	frames := 0
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		// Calculate FPS
		frames++
		if time.Since(lastFrameTime) >= time.Second {
			fmt.Printf("FPS: %d\n", frames)
			frames = 0
			lastFrameTime = time.Now()
		}

		g.Board.ComputeGrid()
		g.Renderer.Clear()
		g.DrawLiveCells()
		g.DrawGrid()
		g.Renderer.Present()
		sdl.Delay(50)
	}
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %s\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Game of life", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(2)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(3)
	}
	defer renderer.Destroy()

	board := gameoflife.NewBoard(screenWidth, screenHeight)

	game := NewGame(renderer, board)
	game.Run()
}
