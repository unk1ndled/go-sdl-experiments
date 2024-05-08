package main

import (
	"fmt"
	"os"
	"time"

	automata "github.com/unk1ndled/nier/automata/logic"
	"github.com/unk1ndled/nier/ds"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	rowNum   = 100
	cellNUm  = 100
	cellSize = 5

	screenWidth  = cellNUm * cellSize
	screenHeight = rowNum * cellSize
)

var (
	rows ds.Queue[[]int8] = *ds.NewQueue[[]int8]()
)

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

	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	quit := false
	// chose rule
	cau := automata.NewRow(cellNUm, automata.Rule{0, 1, 1, 0, 1, 0, 1, 0})

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
		if rows.Length() > rowNum {
			rows.Pop()
		}
		cau.Generate()
		rows.Push(cau.GetGeneration())

		//boxes color
		renderer.SetDrawColor(155, 155, 155, 255)
		//showing cell by cell line by line
		for j, row := range rows.Data {
			for i, cell := range row {
				if cell == 1 {
					renderer.FillRect(&sdl.Rect{X: int32(i * cellSize), Y: int32(j * cellSize), W: cellSize, H: cellSize})
				}
			}
		}

		renderer.Present()
		time.Sleep(40 * time.Millisecond)
	}

}
