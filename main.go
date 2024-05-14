package main

import (
	"fmt"
)

type Board struct {
	size        int
	horizontals [][]int
}

func (b *Board) Init(size int) {
	b.size = size
	zeroes := []int{}

	for i := 0; i < size; i++ {
		zeroes = append(zeroes, 0)
	}

	for i := 0; i < size; i++ {
		b.horizontals = append(b.horizontals, zeroes)
	}
}

func (b *Board) Show() {
	for _, field := range b.horizontals {
		fmt.Println(field)
	}
}

func main() {
	board := Board{}
	board.Init(8)
	board.Show()
}
