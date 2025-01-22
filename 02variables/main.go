package main

import "fmt"

// In go to declare a public variable one must name the 1st letter of variable as Capital
const LoginToken string = "johnny"

func main() {
	var username string = "Narayanan"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	// float32 gives 1st 5 decimal points value
	var smallFloat float32 = 255.33336666
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	// float32 gives more than 5 decimal points value
	var bigFloat float64 = 255.33336666
	fmt.Println(bigFloat)
	fmt.Printf("variable is of type: %T \n", bigFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	// no var style
	// walrus operator :=
	// walrus operator is not allowed outside of the method
	numberOfUser := 3000
	fmt.Println(numberOfUser)
}
