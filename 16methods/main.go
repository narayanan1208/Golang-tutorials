package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Status bool
}

func main() {

	Narayanan := User{"Narayan", 27, true}

	fmt.Printf("%v is %v and active = %v\n", Narayanan.Name, Narayanan.Age, Narayanan.Status)
	Narayanan.getStatus()
	Narayanan.NewName()
}

func (u User) getStatus() {
	fmt.Println("Is user active: ", u.Status)
}

// Changes name with the copy of Narayanan data
func (u User) NewName() {
	u.Name = "Narayanan"
	fmt.Println("New name: ", u.Name)
}
