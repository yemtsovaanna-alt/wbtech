package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(workerID int, dataChannel <-chan int, ctx context.Context, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d shutting down\n", workerID)
			return
		case data, channelOpen := <-dataChannel:
			if !channelOpen {
				return
			}
			fmt.Printf("Worker %d received: %d\n", workerID, data)
		}
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
	ctx, cancel := context.WithCancel(context.Background())
	var waitGroup sync.WaitGroup

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	for workerID := 1; workerID <= numberOfWorkers; workerID++ {
		waitGroup.Add(1)
		go worker(workerID, dataChannel, ctx, &waitGroup)
	}

	go func() {
		<-signalChannel
		fmt.Println("\nReceived interrupt signal, shutting down...")
		cancel()
		close(dataChannel)
	}()

	counter := 0
	for {
		select {
		case <-ctx.Done():
			waitGroup.Wait()
			fmt.Println("All workers stopped")
			return
		default:
			dataChannel <- counter
			counter++
			time.Sleep(100 * time.Millisecond)
		}
	}
}
