package main

import (
	"testing"
)

func TestFirstArray(t *testing.T) {
	array := [...]int{1, 3, 5, 8, 10, 99}
	result, ok := BinarySearch(array[:], 99)
	if ok != nil {
		t.Error("item not found")
	}
	if result != 5 {
		t.Errorf("expected %d but got %d", 5, result)
	}
}

func TestSecondArray(t *testing.T) {
	array := [...]int{1, 3, 5, 8, 10, 99}
	result, ok := BinarySearch(array[:], 1)
	if ok != nil {
		t.Error("item not found")
	}
	if result != 0 {
		t.Errorf("expected %d but got %d", 0, result)
	}
}
