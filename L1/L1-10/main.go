package main

import (
	"fmt"
	"math"
)

func getGroupKey(temperature float64) int {
	return int(math.Floor(temperature/10.0)) * 10
}

func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	
	for _, temperature := range temperatures {
		groupKey := getGroupKey(temperature)
		groups[groupKey] = append(groups[groupKey], temperature)
	}
	
	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	
	fmt.Println("Original temperatures:", temperatures)
	fmt.Println()
	
	groups := groupTemperatures(temperatures)
	
	fmt.Println("Grouped temperatures:")
	for groupKey, values := range groups {
		fmt.Printf("%d: %v\n", groupKey, values)
	}
}
