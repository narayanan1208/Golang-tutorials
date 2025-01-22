package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to rating App")
	fmt.Println("Please our App btwn 1 and 5")

	// bufio is a package for I/O operations
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	// strconv and strings are packages
	numRating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
