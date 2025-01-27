package main

import "fmt"

// Poniter
// when variables are assed, a copy of its data is been shared and due to this sometimes error might occur
// In order to avoid that pointers are used, since it share direct reference to varaible's memory address thereby sharing the actual datas .

func main() {
	fmt.Println("Pointers")

	myNumber := 23
	// & means reference
	var ptr = &myNumber
	fmt.Println("Value of ptr ", ptr)
	fmt.Println("Value of *ptr ", *ptr)

	// Whatever operation is done on pointers, the value of the actual variable is updated instead of the copy of varaible.
	*ptr = *ptr + 2
	fmt.Println("New value is:", myNumber)
}
