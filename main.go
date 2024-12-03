package main


func traverse(board Board) {
	if board.IsWin() {
		board.Show()
		return
	}

	fields := board.FreeFields()

	for _, field := range fields {
		x, y := field[0], field[1]
		copy_board := board.Copy()
		copy_board.Place(x, y)
		traverse(copy_board)
	}
}


func main() {
	board := Board{}
	board.Init(8)
	traverse(board)
}
