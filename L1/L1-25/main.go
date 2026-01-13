package main

import (
	"fmt"
	"runtime"
	"time"
)

func sleepChan(d time.Duration) {
	<-time.After(d)
}

func sleepTimer(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func sleepSelect(d time.Duration) {
	done := make(chan struct{})
	go func() {
		time.AfterFunc(d, func() {
			close(done)
		})
	}()
	<-done
}

func sleepBusy(d time.Duration) {
	end := time.Now().Add(d)
	for time.Now().Before(end) {
		// активное ожидание
	}
}

func sleepBusyYield(d time.Duration) {
	end := time.Now().Add(d)
	for time.Now().Before(end) {
		runtime.Gosched() // уступаем другим горутинам
	}
}

func main() {
	fmt.Println("sleepChan...")
	sleepChan(500 * time.Millisecond)
	fmt.Println("done")

	fmt.Println("sleepTimer...")
	sleepTimer(500 * time.Millisecond)
	fmt.Println("done")

	fmt.Println("sleepSelect...")
	sleepSelect(500 * time.Millisecond)
	fmt.Println("done")

	fmt.Println("sleepBusyYield...")
	sleepBusyYield(500 * time.Millisecond)
	fmt.Println("done")
}
