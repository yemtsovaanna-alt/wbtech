package main

import (
	"fmt"
	"sync"
)

func generateNumbers(numbers []int, outputChannel chan<- int) {
	for _, number := range numbers {
		outputChannel <- number
		fmt.Printf("Generated: %d\n", number)
	}
	close(outputChannel)
}

func multiplyByTwo(inputChannel <-chan int, outputChannel chan<- int) {
	for number := range inputChannel {
		result := number * 2
		outputChannel <- result
		fmt.Printf("Multiplied %d -> %d\n", number, result)
	}
	close(outputChannel)
}

func printResults(inputChannel <-chan int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for result := range inputChannel {
		fmt.Printf("Final result: %d\n", result)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	firstChannel := make(chan int)
	secondChannel := make(chan int)
	
	var waitGroup sync.WaitGroup
	
	go generateNumbers(numbers, firstChannel)
	go multiplyByTwo(firstChannel, secondChannel)
	
	waitGroup.Add(1)
	go printResults(secondChannel, &waitGroup)
	
	waitGroup.Wait()
	fmt.Println("Pipeline completed")
}
