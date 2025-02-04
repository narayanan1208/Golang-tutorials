package main

import (
	"fmt"
)

// A Function cannot be defined inside another function
// Anonymous functions exist
func main() {
	fmt.Println("Functions")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result: ", result)
	proResult, message := proAdder(2, 3, 4, 5, 6)
	fmt.Println("Pro result: ", proResult)
	fmt.Println("pro result message: ", message)
}

// Wrap in paranthesis when two values are returned
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Total"
}

func greeter() {
	fmt.Println("Namaste")
}

// Specify the types of arguments and return
func adder(val1 int, val2 int) int {
	return val1 + val2
}
