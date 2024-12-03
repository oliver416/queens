package main


import "fmt"

func traverse(board Board, solutions *[]Board) {
	if board.IsWin() {
		*solutions = append(*solutions, board)
		return
	}

	fields := board.FreeFields()

	for _, field := range fields {
		x, y := field[0], field[1]
		copy_board := board.Copy()
		copy_board.Place(x, y)
		traverse(copy_board, solutions)
	}
}

func filter(solutions *[]Board) map[string]Board {
	unique := map[string]Board{}

	for _, board := range *solutions {
		hash := board.Hash()

		if _, exists := unique[hash]; !exists {
			unique[hash] = board
		}
	}

	return unique
}


func main() {
	board := Board{}
	board.Init(8)

	solutions := []Board{}
	traverse(board, &solutions)
	filtered := filter(&solutions)

	for _, board := range filtered {
		board.Show()
	}

	fmt.Printf("%d solitions found\n", len(filtered))
}
