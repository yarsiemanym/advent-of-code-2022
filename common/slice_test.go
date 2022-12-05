package common

import (
	"testing"
)

func Test_Slice_Contains_True(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	contains := Contains(slice, 2)

	if !contains {
		t.Error("Expected true but got false.")
	}
}

func Test_Slice_Contains_False(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	contains := Contains(slice, 5)

	if contains {
		t.Error("Expected false but got true.")
	}
}

func Test_Slice_Intersection(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{0, 2, 4, 5}
	intersection := Intersection(slice1, slice2)

	if len(intersection) != 2 {
		t.Errorf("Expected a slice of length 2 but got %d.", len(intersection))
	}
}

func Test_Slice_Union(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{0, 2, 4, 5}
	union := Union(slice1, slice2)

	if len(union) != 6 {
		t.Errorf("Expected a slice of length 6 but got %d.", len(union))
	}
}

func Test_Slice_Exclusive(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{0, 2, 4, 5}
	union := Exclusive(slice1, slice2)

	if len(union) != 4 {
		t.Errorf("Expected a slice of length 4 but got %d.", len(union))
	}
}
