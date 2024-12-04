package main


import "fmt"

func traverse(board Board, solutions map[string]Board, counter *int) {
	*counter += 1

	fields := board.FreeFields()

	for _, field := range fields {
		x, y := field[0], field[1]
		copy_board := board.Copy()
		copy_board.Place(x, y)

		if copy_board.IsWin() {
			hash := copy_board.Hash()

			if _, exists := solutions[hash]; !exists {
				solutions[hash] = copy_board
			}
			return
		}

		traverse(copy_board, solutions, counter)
	}
}

func main() {
	board := Board{}
	board.Init(8)

	solutions := map[string]Board{}
	counter := 0
	traverse(board, solutions, &counter)

	for _, board := range solutions {
		board.Show()
	}

	fmt.Printf("%d solitions found, %d iterations\n", len(solutions), counter)
}
