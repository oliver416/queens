package tests

import (
	"golang.org/x/exp/slices"
	"queens/core/array"
	"testing"
)

func TestMoveArrayElement(t *testing.T) {
	// TODO: unify the tests
	arr := []int{4, 1, 5, 2, 3}
	expected := []int{4, 2, 1, 5, 3}

	result := array.MoveArrayElement(arr, 3, 1)

	if !slices.Equal(result, expected) {
		t.Error(result)
	}
}
