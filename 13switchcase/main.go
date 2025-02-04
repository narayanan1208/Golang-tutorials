package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case")

	// rand.Seed(time.Now().UnixNano()) deprecated
	// seed := time.Now().UnixNano()
	// src := rand.NewSource(seed)
	// r := rand.New(src)
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and start")
	case 2:
		fmt.Println("Move 2 spot")
	case 3:
		fmt.Println("Move 3 spot")
		fallthrough // similar to continue statement
	case 4:
		fmt.Println("Move 4 spot")
	case 5:
		fmt.Println("Move 5 spot")
	case 6:
		fmt.Println("Move 6 spot and roll again")
	default:
		fmt.Print("Unknown roll")
	}
}
