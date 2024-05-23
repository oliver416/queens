package main

import (
	"fmt"
	"queens/array"
)

func main() {
	arr := []int{3, 1, 2, 2, 3, 1, 4, 2}

	fmt.Println(array.RemoveArrayElementInplace(arr, 2))
    fmt.Println(array.SelectionSort(arr))
    fmt.Println(array.BubbleSort(arr))
}
