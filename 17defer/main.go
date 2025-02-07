package main

import "fmt"

// defer follows last-in first out. when defer is used before a line, it gets executed last.
func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
