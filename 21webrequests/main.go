package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web request methods")
	// PerformGetRequest()
	PerformPostFormRequest()
	// PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Response status: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	// fmt.Println("Content: ", string(content))
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count: ", byteCount)
	fmt.Println("Response content: ", responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform": "lco.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkNilErr(err)
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "Narayanan")
	data.Add("lastname", "S")
	data.Add("email", "narayanans@gmail.com")

	response, err := http.PostForm(myurl, data)
	checkNilErr(err)
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
