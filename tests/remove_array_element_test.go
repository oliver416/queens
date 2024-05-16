package tests

import (
	"golang.org/x/exp/slices"
	"testing"
    "queens/array"
)

func TestRemoveArrayElementInplace(t *testing.T) {
	arr := []int{3, 1, 2, 2, 3, 1, 4, 2}
    expected := []int{3, 1, 3, 1, 4}

	result := array.RemoveArrayElementInplace(arr, 2)

	if !slices.Equal(result, expected) {
		t.Error(result)
	}
}
