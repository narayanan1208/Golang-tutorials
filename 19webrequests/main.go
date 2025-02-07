package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("Web requests")
	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Response type: %T\n", response)
	defer response.Body.Close() // developer's responsibility to close the response

	databyte, err := io.ReadAll(response.Body)
	checkNilErr(err)
	content := string(databyte)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
