package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type MutexCounter struct {
	mu    sync.Mutex
	value int
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

type AtomicCounter struct {
	value atomic.Int64
}

func (c *AtomicCounter) Inc() {
	c.value.Add(1)
}

func (c *AtomicCounter) Value() int64 {
	return c.value.Load()
}

func main() {
	mutexCounter := &MutexCounter{}
	atomicCounter := &AtomicCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			mutexCounter.Inc()
		}()
		go func() {
			defer wg.Done()
			atomicCounter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("MutexCounter:", mutexCounter.Value())   // 100
	fmt.Println("AtomicCounter:", atomicCounter.Value()) // 100
}
