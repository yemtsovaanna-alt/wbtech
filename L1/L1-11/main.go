package main

import "fmt"

func intersectSlices(firstSlice, secondSlice []int) []int {
	elementSet := make(map[int]bool)
	intersection := []int{}
	
	for _, element := range firstSlice {
		elementSet[element] = true
	}
	
	visited := make(map[int]bool)
	for _, element := range secondSlice {
		if elementSet[element] && !visited[element] {
			intersection = append(intersection, element)
			visited[element] = true
		}
	}
	
	return intersection
}

func intersectSlicesGeneric[T comparable](firstSlice, secondSlice []T) []T {
	elementSet := make(map[T]bool)
	intersection := []T{}
	
	for _, element := range firstSlice {
		elementSet[element] = true
	}
	
	visited := make(map[T]bool)
	for _, element := range secondSlice {
		if elementSet[element] && !visited[element] {
			intersection = append(intersection, element)
			visited[element] = true
		}
	}
	
	return intersection
}

func main() {
	fmt.Println("=== Integer slices ===")
	firstSet := []int{1, 2, 3}
	secondSet := []int{2, 3, 4}
	
	fmt.Printf("A = %v\n", firstSet)
	fmt.Printf("B = %v\n", secondSet)
	fmt.Printf("Intersection = %v\n", intersectSlices(firstSet, secondSet))
	fmt.Println()
	
	fmt.Println("=== With duplicates ===")
	thirdSet := []int{1, 2, 2, 3, 3, 4}
	fourthSet := []int{2, 3, 3, 4, 5, 5}
	
	fmt.Printf("A = %v\n", thirdSet)
	fmt.Printf("B = %v\n", fourthSet)
	fmt.Printf("Intersection = %v\n", intersectSlices(thirdSet, fourthSet))
	fmt.Println()
	
	fmt.Println("=== String slices (generic) ===")
	firstStringSet := []string{"apple", "banana", "cherry"}
	secondStringSet := []string{"banana", "cherry", "date"}
	
	fmt.Printf("A = %v\n", firstStringSet)
	fmt.Printf("B = %v\n", secondStringSet)
	fmt.Printf("Intersection = %v\n", intersectSlicesGeneric(firstStringSet, secondStringSet))
	fmt.Println()
	
	fmt.Println("=== Empty intersection ===")
	fifthSet := []int{1, 2, 3}
	sixthSet := []int{4, 5, 6}
	
	fmt.Printf("A = %v\n", fifthSet)
	fmt.Printf("B = %v\n", sixthSet)
	fmt.Printf("Intersection = %v\n", intersectSlices(fifthSet, sixthSet))
}
