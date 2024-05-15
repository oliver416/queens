package main

import (
	"fmt"
)

// TODO: poor naming
// TODO: add a documentation
func remove(nums []int, item int) []int {
	index := 0

	for _, value := range nums {
		if value != item {
			nums[index] = value
			index++
		}
	}
	return nums[:index]
}

func main() {
    // TODO: cover with test
	array := []int{3, 1, 2, 2, 3, 1, 4, 2}
	result := remove(array, 2)
	fmt.Println(result)
}
