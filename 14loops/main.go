package main

import "fmt"

func main() {
	fmt.Println("Loops")
	days := []string{"SUN", "SAT", "FRI", "THU", "WED", "TUE", "MON"}

	fmt.Println(days)

	// Types of loops

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// i gives the index
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	value := 1
	for value < 10 {

		if value == 5 {
			break
			// value++
			// continue
			// goto lco
		}
		fmt.Println("Value is :", value)
		value++ // only ending with ++
	}

	// lco:
	// 	fmt.Println("Value is 5")

}
