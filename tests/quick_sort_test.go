package tests

import (
	"golang.org/x/exp/slices"
	"queens/core/array"
	"testing"
)

func TestDivideArray(t *testing.T) {
	arr := []int{3, 1, 2, 2, 3, 1, 4, 2}
	expected := []int{1, 2, 2, 1, 2, 3, 4, 3}

	result, divider := array.DivideArray(arr)

	if !slices.Equal(result, expected) {
		t.Error(result)
	}

	if divider != 5 {
		t.Error(divider)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 1, 2, 2, 3, 1, 4, 2}
	expected := []int{1, 1, 2, 2, 2, 3, 3, 4}

	result := array.QuickSort(arr)

	if !slices.Equal(result, expected) {
		t.Error(result)
	}
}
