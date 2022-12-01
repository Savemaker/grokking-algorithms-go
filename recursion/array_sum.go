package main

func ArraySum(arr []int) int {
	len := len(arr)
	mid := len / 2
	if len == 1 {
		return arr[0]
	}
	return ArraySum(arr[:mid]) + ArraySum(arr[mid:])
}
