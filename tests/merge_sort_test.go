package tests

import (
	"golang.org/x/exp/slices"
	"testing"
    "queens/array"
)

func TestMerge(t *testing.T) {
    arr1 := []int{1, 2, 2, 3, 4}
    arr2 := []int{1, 2, 3}
    expected := []int{1, 1, 2, 2, 2, 3, 3, 4}

	result := array.Merge(arr1, arr2)

	if !slices.Equal(result, expected) {
		t.Error(result)
	}
}


func TestMergeSort(t *testing.T) {
	arr := []int{4, 1, 5, 2, 3}
    expected := []int{1, 2, 3, 4, 5}

	result := array.MergeSort(arr)

	if !slices.Equal(result, expected) {
		t.Error(result)
	}
}

