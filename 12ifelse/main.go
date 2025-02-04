package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("If else")

	// loginCount := 23
	fmt.Println("Enter the count: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	loginCount, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}

	var result string

	if loginCount < 10 {
		result = "Lesser"
	} else if loginCount > 10 {
		result = "Greater"
	} else {
		result = "Equal"
	}

	fmt.Println(result)

	// initialization and value assignment in one line
	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("greater than 10")
	}

}
