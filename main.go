package main

func main() {
	board := Board{}
	board.Init(8)
	// TODO: cover it with tests
	// board.Place(3,4)
	board.Place(7, 2)
	board.Show()
}
