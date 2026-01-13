package main

import "fmt"

// удаление с сохранением порядка (без утечки памяти)
func removeOrdered[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		return slice
	}
	copy(slice[i:], slice[i+1:])
	var zero T
	slice[len(slice)-1] = zero // обнуляем последний элемент для GC
	return slice[:len(slice)-1]
}

// удаление без сохранения порядка (быстрее, O(1))
func removeUnordered[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		return slice
	}
	slice[i] = slice[len(slice)-1]
	var zero T
	slice[len(slice)-1] = zero
	return slice[:len(slice)-1]
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s1 = removeOrdered(s1, 2)
	fmt.Println(s1) // [1 2 4 5]

	s2 := []int{1, 2, 3, 4, 5}
	s2 = removeUnordered(s2, 1)
	fmt.Println(s2) // [1 5 3 4]

	// указатели (важно обнулять для GC)
	type User struct{ Name string }
	users := []*User{{"Alice"}, {"Bob"}, {"Charlie"}}
	users = removeOrdered(users, 1)
	fmt.Println(users[0].Name, users[1].Name) // Alice Charlie
}
