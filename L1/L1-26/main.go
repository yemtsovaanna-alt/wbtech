package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range strings.ToLower(s) {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd")) // true

	fmt.Println(isUnique("abCdefAaf")) // false

	fmt.Println(isUnique("aabcd")) // false

	fmt.Println(isUnique("привет"))  // true
	fmt.Println(isUnique("приветр")) // false
}
