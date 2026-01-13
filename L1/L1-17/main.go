package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}

	fmt.Println(binarySearch(arr, 7)) // 3

	fmt.Println(binarySearch(arr, 13)) // 6

	fmt.Println(binarySearch(arr, 4)) // -1
}
