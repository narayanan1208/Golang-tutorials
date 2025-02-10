package main

import (
	"fmt"
	"time"
)

func main() {
	// go runs in a separate goroutine and execution is asynchronous
	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
