package main

import (
	"fmt"
	"net/http"
	"sync"
)

// A Mutex is used to prevent race conditions by ensuring that only one goroutine can access a critical section of code at a time.
// This is useful when multiple goroutines need to read and write shared resources.

// A WaitGroup is used to wait for a collection of goroutines to finish executing.
// You increment the counter with Add(), decrement it with Done(), and block until the counter becomes zero with Wait().
var signals = []string{"test"}
var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func main() {
	websitelist := []string{
		"https://github.com",
		"https://google.com",
		"https://go.dev",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
