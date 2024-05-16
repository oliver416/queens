package main

import (
	"fmt"
    "queens/array"
)

func main() {
    // TODO: cover with test
	arr := []int{3, 1, 2, 2, 3, 1, 4, 2}
	result := array.RemoveArrayElementInplace(arr, 2)
	fmt.Println(result)
}
