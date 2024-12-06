package tests

import (
	"golang.org/x/exp/slices"
	"queens/core/array"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 1, 5, 2, 3}
	expected := []int{1, 2, 3, 4, 5}

	result := array.SelectionSort(arr)

	if !slices.Equal(result, expected) {
		t.Error(result)
	}
}
