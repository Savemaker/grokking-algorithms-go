package main

import "testing"

func TestArraySum(t *testing.T) {
	arr := [...]int{1, -2, 69, 4, 7}

	res := ArraySum(arr[:])

	if res != 79 {
		t.Errorf("expected 79 but got %d", res)
	}

}
