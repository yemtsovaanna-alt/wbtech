package main

import "fmt"

func createSet(strings []string) map[string]struct{} {
	uniqueSet := make(map[string]struct{})
	
	for _, str := range strings {
		uniqueSet[str] = struct{}{}
	}
	
	return uniqueSet
}

func setToSlice(set map[string]struct{}) []string {
	result := make([]string, 0, len(set))
	
	for key := range set {
		result = append(result, key)
	}
	
	return result
}

func createSetWithBool(strings []string) map[string]bool {
	uniqueSet := make(map[string]bool)
	
	for _, str := range strings {
		uniqueSet[str] = true
	}
	
	return uniqueSet
}

func getUniqueStrings(strings []string) []string {
	seen := make(map[string]bool)
	unique := []string{}
	
	for _, str := range strings {
		if !seen[str] {
			unique = append(unique, str)
			seen[str] = true
		}
	}
	
	return unique
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	
	fmt.Println("=== Original sequence ===")
	fmt.Printf("Words: %v\n", words)
	fmt.Println()
	
	fmt.Println("=== Set with struct{} ===")
	setWithStruct := createSet(words)
	fmt.Printf("Set (map): %v\n", setWithStruct)
	fmt.Printf("Set (slice): %v\n", setToSlice(setWithStruct))
	fmt.Printf("Size: %d\n", len(setWithStruct))
	fmt.Println()
	
	fmt.Println("=== Set with bool ===")
	setWithBool := createSetWithBool(words)
	fmt.Printf("Set: %v\n", setWithBool)
	fmt.Printf("Size: %d\n", len(setWithBool))
	fmt.Println()
	
	fmt.Println("=== Unique slice (preserving order) ===")
	uniqueWords := getUniqueStrings(words)
	fmt.Printf("Unique: %v\n", uniqueWords)
	fmt.Println()
	
	fmt.Println("=== Set operations ===")
	testSet := createSet(words)
	fmt.Printf("Contains 'cat': %v\n", func() bool {
		_, exists := testSet["cat"]
		return exists
	}())
	fmt.Printf("Contains 'bird': %v\n", func() bool {
		_, exists := testSet["bird"]
		return exists
	}())
}
