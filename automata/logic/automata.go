package automata

import (
	"fmt"
	"time"
)

type Rule [8]int8

type Row struct {
	generation []int8
}

func NewRow(length int) *Row {
	row := Row{generation: make([]int8, length)}
	row.generation[length/2] = 1
	return &row

}

func (row *Row) GetGeneration() []int8 {
	return row.generation
}

func (row *Row) Generate(rule Rule) {
	size := len(row.generation)
	nextgen := make([]int8, size)
	for i := 0; i < size; i++ {
		value := Compute(row.generation[(i-1+size)%size], row.generation[i], row.generation[(i+1)%size], rule)
		nextgen[i] = value

	}
	row.generation = nextgen
}

func Compute(i1, i2, i3 int8, rule Rule) int8 {
	index := i1*4 + i2*2 + i3
	return rule[index]
}

func StartOneD() {
	rule80 := Rule{0, 1, 1, 0, 1, 0, 1, 0}

	row := NewRow(110)
	var rowval []int8

	for {
		row.Generate(rule80)
		rowval = row.GetGeneration()
		fmt.Println(rowval)
		time.Sleep(50 * time.Millisecond)
	}
}
