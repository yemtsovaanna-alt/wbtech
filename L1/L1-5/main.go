package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func sender(dataChannel chan<- int, doneChannel <-chan struct{}) {
	counter := 0
	for {
		select {
		case <-doneChannel:
			close(dataChannel)
			return
		case dataChannel <- counter:
			fmt.Printf("Sent: %d\n", counter)
			counter++
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func receiver(dataChannel <-chan int, doneChannel <-chan struct{}) {
	for {
		select {
		case <-doneChannel:
			return
		case value, channelOpen := <-dataChannel:
			if !channelOpen {
				return
			}
			fmt.Printf("Received: %d\n", value)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <seconds>")
		return
	}

	seconds, err := strconv.Atoi(os.Args[1])
	if err != nil || seconds < 1 {
		fmt.Println("Invalid number of seconds")
		return
	}

	dataChannel := make(chan int)
	doneChannel := make(chan struct{})

	go sender(dataChannel, doneChannel)
	go receiver(dataChannel, doneChannel)

	<-time.After(time.Duration(seconds) * time.Second)
	
	fmt.Println("\nTime limit reached, shutting down...")
	close(doneChannel)
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Program finished")
}
