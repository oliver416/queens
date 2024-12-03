## How to check a board
```go
package main

import "fmt"


func main() {
	board := Board{}
	board.Init(8)
	board.Place(3,4)
	board.Place(7, 2)
	board.Show()

	sum := board.Sum()
	fmt.Println(sum)

	fields := board.FreeFields()
	fmt.Println(fields)
}

```