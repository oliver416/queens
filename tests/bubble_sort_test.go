package tests

import (
	"golang.org/x/exp/slices"
	"testing"
    "queens/array"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 1, 5, 2, 3}
    expected := []int{1, 2, 3, 4, 5}

	result := array.BubbleSort(arr)

	if !slices.Equal(result, expected) {
		t.Error(result)
	}
}
