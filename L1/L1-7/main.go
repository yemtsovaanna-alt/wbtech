package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMapWithMutex struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMapWithMutex() *SafeMapWithMutex {
	return &SafeMapWithMutex{
		data: make(map[string]int),
	}
}

func (sm *SafeMapWithMutex) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMapWithMutex) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

func (sm *SafeMapWithMutex) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SafeMapWithMutex) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

type SafeMapWithSyncMap struct {
	data sync.Map
}

func NewSafeMapWithSyncMap() *SafeMapWithSyncMap {
	return &SafeMapWithSyncMap{}
}

func (sm *SafeMapWithSyncMap) Set(key string, value int) {
	sm.data.Store(key, value)
}

func (sm *SafeMapWithSyncMap) Get(key string) (int, bool) {
	value, exists := sm.data.Load(key)
	if !exists {
		return 0, false
	}
	return value.(int), true
}

func (sm *SafeMapWithSyncMap) Delete(key string) {
	sm.data.Delete(key)
}

func (sm *SafeMapWithSyncMap) Len() int {
	length := 0
	sm.data.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	return length
}

func testSafeMapWithMutex() {
	fmt.Println("=== Testing SafeMap with Mutex ===")
	safeMap := NewSafeMapWithMutex()
	var waitGroup sync.WaitGroup
	numberOfGoroutines := 100

	for goroutineID := 0; goroutineID < numberOfGoroutines; goroutineID++ {
		waitGroup.Add(1)
		go func(id int) {
			defer waitGroup.Done()
			key := fmt.Sprintf("key_%d", id)
			safeMap.Set(key, id*10)
			time.Sleep(time.Millisecond)
			
			value, exists := safeMap.Get(key)
			if exists {
				fmt.Printf("Goroutine %d: %s = %d\n", id, key, value)
			}
		}(goroutineID)
	}

	waitGroup.Wait()
	fmt.Printf("Total entries: %d\n\n", safeMap.Len())
}

func testSafeMapWithSyncMap() {
	fmt.Println("=== Testing SafeMap with sync.Map ===")
	safeMap := NewSafeMapWithSyncMap()
	var waitGroup sync.WaitGroup
	numberOfGoroutines := 100

	for goroutineID := 0; goroutineID < numberOfGoroutines; goroutineID++ {
		waitGroup.Add(1)
		go func(id int) {
			defer waitGroup.Done()
			key := fmt.Sprintf("key_%d", id)
			safeMap.Set(key, id*10)
			time.Sleep(time.Millisecond)
			
			value, exists := safeMap.Get(key)
			if exists {
				fmt.Printf("Goroutine %d: %s = %d\n", id, key, value)
			}
		}(goroutineID)
	}

	waitGroup.Wait()
	fmt.Printf("Total entries: %d\n\n", safeMap.Len())
}

func testConcurrentReadWrite() {
	fmt.Println("=== Testing Concurrent Read/Write ===")
	safeMap := NewSafeMapWithMutex()
	var waitGroup sync.WaitGroup
	numberOfWriters := 50
	numberOfReaders := 50

	for writerID := 0; writerID < numberOfWriters; writerID++ {
		waitGroup.Add(1)
		go func(id int) {
			defer waitGroup.Done()
			for iteration := 0; iteration < 10; iteration++ {
				key := fmt.Sprintf("writer_%d_%d", id, iteration)
				safeMap.Set(key, id*100+iteration)
			}
		}(writerID)
	}

	for readerID := 0; readerID < numberOfReaders; readerID++ {
		waitGroup.Add(1)
		go func(id int) {
			defer waitGroup.Done()
			for iteration := 0; iteration < 10; iteration++ {
				key := fmt.Sprintf("writer_%d_%d", id%numberOfWriters, iteration)
				if value, exists := safeMap.Get(key); exists {
					fmt.Printf("Reader %d read: %s = %d\n", id, key, value)
				}
				time.Sleep(time.Millisecond)
			}
		}(readerID)
	}

	waitGroup.Wait()
	fmt.Printf("Final map size: %d\n\n", safeMap.Len())
}

func main() {
	testSafeMapWithMutex()
	testSafeMapWithSyncMap()
	testConcurrentReadWrite()
	
	fmt.Println("All tests completed successfully")
	fmt.Println("Run with: go run -race main.go")
}
