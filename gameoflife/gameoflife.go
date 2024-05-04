package gameoflife

import (
	"math/rand"
)

type Grid [][]int16

type Board struct {
	grid      *Grid
	Livecells *[][2]int16
}

func (b *Board) GetGrid() Grid {
	return (*b.grid)
}
func (b *Board) ComputeGrid() {
	grid := b.GetGrid()
	height := len(grid)
	width := len(grid[0])
	b.Livecells = &[][2]int16{}
	temp := NewClearGrid(width, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			isAlive := grid.ComputeNeighbours(i, j)
			(*temp)[i][j] = isAlive
			if isAlive == 1 {
				*b.Livecells = append((*b.Livecells), [2]int16{int16(i), int16(j)})
			}
		}
	}
	b.grid = temp
}

func (g *Grid) ComputeNeighbours(i, j int) int16 {
	size := len(*g)
	center := (*g)[i][j]
	var neighbors int16

	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			xIndex := (x + size) % size
			yIndex := (y + size) % size
			neighbors += (*g)[xIndex][yIndex]
		}
	}
	neighbors -= center

	if center == 1 && neighbors < 2 {
		return 0 //Loner
	} else if center == 1 && neighbors > 3 {
		return 0 //Doomer
	} else if center == 0 && neighbors == 3 {
		return 1 //Bloomer
	} else {
		//Sleeper
		return center
	}
}

func NewBoard(width, height int) *Board {
	cells := [][2]int16{}
	return &Board{grid: NewGrid(width, height, &cells), Livecells: &cells}
}

func NewClearGrid(width, height int) *Grid {
	grid := make([][]int16, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int16, width)
	}
	return (*Grid)(&grid)
}

func NewGrid(width, height int, livecells *[][2]int16) *Grid {
	grid := make([][]int16, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int16, width)
		for j := 0; j < width; j++ {
			value := int16(rand.Intn(2))
			grid[i][j] = value
			if value == 1 {
				(*livecells) = append((*livecells), [2]int16{int16(i), int16(j)})
			}

		}
	}
	return (*Grid)(&grid)
}
