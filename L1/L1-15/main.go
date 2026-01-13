package main

import "strings"

func createHugeString(int) string {
	return ""
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}
