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
		copy_zeroes := make([]int, len(zeroes))
		copy(copy_zeroes, zeroes)
		b.horizontals = append(b.horizontals, copy_zeroes)
	}
}

func (b *Board) Show() {
	for _, field := range b.horizontals {
		fmt.Println(field)
	}
}

func (b *Board) Place(x int, y int) {
	// TODO: what if x or y < 0??
	if x < b.size && y < b.size {
		for i := 0; i < b.size; i++ {
			b.horizontals[x][i] = 1
			b.horizontals[i][y] = 1
		}

		for i, j := x, y; i >= 0 && j >= 0; {
			b.horizontals[i][j] = 1
			i--
			j--
		}

		for i, j := x, y; i < b.size && j >= 0; {
			b.horizontals[i][j] = 1
			i++
			j--
		}

		for i, j := x, y; i >= 0 && j < b.size; {
			b.horizontals[i][j] = 1
			i--
			j++
		}

		for i, j := x, y; i <= b.size-1 && j <= b.size-1; {
			b.horizontals[i][j] = 1
			i++
			j++
		}

		b.horizontals[x][y] = 8
	}
}
