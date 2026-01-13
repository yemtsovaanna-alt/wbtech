package main

import "fmt"

func reverseWords(s string) string {
	runes := []rune(s)

	reverse(runes, 0, len(runes)-1) // разворачиваем строку целиком

	// и потом разворачиваем слова
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func reverse(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

func main() {
	fmt.Println(reverseWords("snow dog sun")) // sun dog snow

	fmt.Println(reverseWords("привет мир всем")) // всем мир привет

	fmt.Println(reverseWords("hello")) // hello
}
