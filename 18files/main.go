package main

import (
	"fmt"
	"io"
	"os"
)

// for creating other file formats there are inbuild packages in go.

func main() {
	fmt.Println("Files")
	content := "Data writing"
	file, err := os.Create("./CreatedFile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("Length: ", length)
	defer file.Close()
	readFile("./CreatedFile.txt")
}

func readFile(filname string) {
	databyte, err := os.ReadFile(filname)
	checkNilErr(err)
	// type conversion is needed bcoz databyte is not readable
	fmt.Println("Data inside file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
