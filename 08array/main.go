package main

// go mod init command creates a go.mod file to track your code's dependencies
import "fmt"

func main() {
	fmt.Println("Arrays")

	// Array size is to be mentioned mandatory
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Orange"

	fmt.Println("Fruit list: ", fruitList)
	fmt.Println("Len of fruit list: ", len(fruitList))

	var vegList = [3]string{"Onion", "Beans", "Raddish"}
	fmt.Println("Veg list: ", vegList)
}
