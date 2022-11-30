package main

import "errors"

func BinarySearch(arr []int, query int) (int, error) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (high + low) / 2
		if arr[mid] == query {
			return mid, nil
		} else if arr[mid] < query {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, errors.New("item does not exist")
}
