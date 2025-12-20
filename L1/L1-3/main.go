package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(workerID int, dataChannel <-chan int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for data := range dataChannel {
		fmt.Printf("Worker %d received: %d\n", workerID, data)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <number_of_workers>")
		return
	}

	numberOfWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numberOfWorkers < 1 {
		fmt.Println("Invalid number of workers")
		return
	}

	dataChannel := make(chan int)
	var waitGroup sync.WaitGroup

	for workerID := 1; workerID <= numberOfWorkers; workerID++ {
		waitGroup.Add(1)
		go worker(workerID, dataChannel, &waitGroup)
	}

	for counter := 0; counter < 20; counter++ {
		dataChannel <- counter
		time.Sleep(100 * time.Millisecond)
	}

	close(dataChannel)
	waitGroup.Wait()
}
