package main

import (
	"fmt"
	"sync"
)

// go run --race main.go cmd shows the underlying process
// This program does not flow in order

// An RWMutex (Read-Write Mutex) is a synchronization primitive provided by the sync package.
// It allows multiple goroutines to read from a shared resource simultaneously, but only one goroutine to write to the resource at a time. This ensures data consistency while optimizing performance for read-heavy workloads.
// RLock(): Acquires the read lock (shared). Multiple readers can acquire this lock simultaneously.
// RUnlock(): Releases the read lock.

// Lock(): Acquires the write lock (exclusive). Blocks other readers and writers.
// Unlock(): Releases the write lock.

func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{} // to prevent race conditions

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Read Write Mutex")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
