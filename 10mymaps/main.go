package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)
	languages["JAVA"] = "Java"
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("languages: ", languages["JS"])

	// delete
	delete(languages, "PY")
	fmt.Println("List of all languages: ", languages)

	// loops in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
