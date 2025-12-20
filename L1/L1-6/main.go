package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func stopByCondition() {
	fmt.Println("=== Stop by condition ===")
	var waitGroup sync.WaitGroup
	stopFlag := false

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for !stopFlag {
			fmt.Println("Working...")
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("Stopped by condition")
	}()

	time.Sleep(1 * time.Second)
	stopFlag = true
	waitGroup.Wait()
	fmt.Println()
}

func stopByChannel() {
	fmt.Println("=== Stop by channel ===")
	var waitGroup sync.WaitGroup
	stopChannel := make(chan struct{})

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			select {
			case <-stopChannel:
				fmt.Println("Stopped by channel")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(stopChannel)
	waitGroup.Wait()
	fmt.Println()
}

func stopByContext() {
	fmt.Println("=== Stop by context ===")
	var waitGroup sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopped by context:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()
	waitGroup.Wait()
	fmt.Println()
}

func stopByContextWithTimeout() {
	fmt.Println("=== Stop by context with timeout ===")
	var waitGroup sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopped by timeout:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	waitGroup.Wait()
	fmt.Println()
}

func stopByContextWithDeadline() {
	fmt.Println("=== Stop by context with deadline ===")
	var waitGroup sync.WaitGroup
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopped by deadline:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	waitGroup.Wait()
	fmt.Println()
}

func stopByReturn() {
	fmt.Println("=== Stop by return ===")
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for iteration := 0; iteration < 5; iteration++ {
			fmt.Println("Working iteration:", iteration)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("Stopped by return after iterations")
		return
	}()

	waitGroup.Wait()
	fmt.Println()
}

func stopByGoexit() {
	fmt.Println("=== Stop by runtime.Goexit ===")
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		defer fmt.Println("Deferred function executed before Goexit")
		
		for iteration := 0; iteration < 3; iteration++ {
			fmt.Println("Working iteration:", iteration)
			time.Sleep(200 * time.Millisecond)
		}
		
		fmt.Println("Calling runtime.Goexit")
		runtime.Goexit()
		fmt.Println("This will never be printed")
	}()

	waitGroup.Wait()
	fmt.Println()
}

func stopByPanic() {
	fmt.Println("=== Stop by panic with recover ===")
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		defer func() {
			if recoverValue := recover(); recoverValue != nil {
				fmt.Println("Recovered from panic:", recoverValue)
			}
		}()

		for iteration := 0; iteration < 3; iteration++ {
			fmt.Println("Working iteration:", iteration)
			time.Sleep(200 * time.Millisecond)
		}

		panic("intentional panic to stop goroutine")
	}()

	waitGroup.Wait()
	fmt.Println()
}

func stopByCombinedChannels() {
	fmt.Println("=== Stop by combined channels (done + data) ===")
	var waitGroup sync.WaitGroup
	dataChannel := make(chan int)
	doneChannel := make(chan struct{})

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			select {
			case <-doneChannel:
				fmt.Println("Stopped by done channel")
				return
			case data := <-dataChannel:
				fmt.Println("Received data:", data)
			}
		}
	}()

	dataChannel <- 1
	dataChannel <- 2
	time.Sleep(500 * time.Millisecond)
	close(doneChannel)
	waitGroup.Wait()
	fmt.Println()
}

func main() {
	stopByCondition()
	stopByChannel()
	stopByContext()
	stopByContextWithTimeout()
	stopByContextWithDeadline()
	stopByReturn()
	stopByGoexit()
	stopByPanic()
	stopByCombinedChannels()

	fmt.Println("All examples completed")
}
