package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Banana", "Orange"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Dates")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 100
	highScores[1] = 200
	highScores[2] = 300
	highScores[3] = 400

	fmt.Println("High scores: ", highScores)

	// append allocates the new values into slices
	highScores = append(highScores, 500, 600)
	fmt.Println("High scores after append: ", highScores)

	// sort function used with int
	sort.Ints(highScores)
	fmt.Println("Sorted high scores: ", highScores)
	// boolean return function
	fmt.Println("Is sorted or not: ", sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index
	var alphabets = []string{"a", "b", "c", "d"}
	fmt.Println(alphabets)
	var index int = 2
	alphabets = append(alphabets[:index], alphabets[index+1:]...)
	fmt.Println("After append: ", alphabets)
}
