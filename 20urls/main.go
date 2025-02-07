package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com/search?q=beckham"

func main() {
	fmt.Println("Handling urls")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qParams)

	for _, val := range qParams {
		fmt.Println("Param is ", val)
	}

	partsOfUrl := &url.URL{Scheme: "https", Host: "www.google.com", Path: "/search", RawPath: "beckham"}

	createdURL := partsOfUrl.String()
	fmt.Println(createdURL)
}
