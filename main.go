package main

import (
	"fmt"
	"queens/array"
)

func main() {
	arr := []int{3, 1, 2, 2, 3, 1, 4, 2}

    fmt.Println(arr)
    
    // TODO: fix a slice returning bug
	//fmt.Println(array.RemoveArrayElementInplace(arr, 2))

    fmt.Println(array.MoveArrayElement(arr, 6, 1))

    fmt.Println("Sorting algorithms:")
	arr = []int{3, 1, 2, 2, 3, 1, 4, 2}
    fmt.Println(array.InsertionSort(arr))

	arr = []int{3, 1, 2, 2, 3, 1, 4, 2}
    fmt.Println(array.SelectionSort(arr))

	arr = []int{3, 1, 2, 2, 3, 1, 4, 2}
    fmt.Println(array.BubbleSort(arr))

    arr1 := []int{1, 2, 2, 3, 4}
    arr2 := []int{1, 2, 3}
    fmt.Println(array.Merge(arr1, arr2))

	arr = []int{3, 1, 2, 2, 3, 1, 4, 2}
    fmt.Println(array.MergeSort(arr))
}
