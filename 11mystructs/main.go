package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// No inheritance in golang; No super or parent

	Narayanan := User{"Narayanan", "narayanan6647@gmail.com", true, 27}
	fmt.Println("Narayanan: ", Narayanan)
	fmt.Printf("Details are: %+v\n", Narayanan)
	fmt.Printf("Name is %v and email is %v.", Narayanan.Name, Narayanan.Email)
}

// name should start with Capital letter
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
