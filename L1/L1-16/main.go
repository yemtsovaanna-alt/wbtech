package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]
	var left, middle, right []int
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}
	result := append(quickSort(left), middle...)
	return append(result, quickSort(right)...)
}

func main() {
	fmt.Println(quickSort([]int{3, 1, 4, 1, 5, 9, 2, 6})) // [1 1 2 3 4 5 6 9]

	fmt.Println(quickSort([]int{-5, 10, -3, 0, 7, -1})) // [-5 -3 -1 0 7 10]

	fmt.Println(quickSort([]int{9, 7, 5, 3, 1})) // [1 3 5 7 9]
}
