package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var waitGroup sync.WaitGroup

	// в новой версии го variable per itaration, поэтому не нужно делать ничего лишнего с переменной
	for _, number := range numbers {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			squaredValue := number * number
			fmt.Println(squaredValue)
		}()
	}

	waitGroup.Wait()
}
